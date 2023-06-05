package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"favor-notify/internal"
	"favor-notify/internal/conf"
	"favor-notify/internal/routers"
	"favor-notify/pkg/debug"
	"favor-notify/pkg/util"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

var (
	noDefaultFeatures bool
	features          suites
)

type suites []string

func (s *suites) String() string {
	return strings.Join(*s, ",")
}

func (s *suites) Set(value string) error {
	for _, item := range strings.Split(value, ",") {
		*s = append(*s, strings.TrimSpace(item))
	}
	return nil
}

func init() {
	conf.Initialize(features, noDefaultFeatures)
	internal.Initialize()
}

func main() {

	gin.SetMode(conf.ServerSetting.RunMode)

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           conf.ServerSetting.HttpIp + ":" + conf.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    conf.ServerSetting.ReadTimeout,
		WriteTimeout:   conf.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	util.PrintHelloBanner(debug.VersionInfo())
	fmt.Fprintf(color.Output, "favor notify service listen on %s\n",
		color.GreenString(fmt.Sprintf("https://%s:%s", conf.ServerSetting.HttpIp, conf.ServerSetting.HttpPort)),
	)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("app stoped: %s", err)
		}
		sigs <- syscall.SIGTERM
	}()
	select {
	case <-sigs:
	}
	s.Shutdown(context.TODO())
}
