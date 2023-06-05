package routers

import (
	"favor-notify/internal/routers/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	r := e.Group("/v1")
	r.POST("/push/notify", api.PushNotify)
	r.POST("/push/notify/dao", api.PushNotifyDao)
	r.POST("/push/notify/sys", api.PushNotifySys)
	e.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"code": "404",
			"msg":  "Not Found",
		})
	})
	e.NoMethod(func(context *gin.Context) {
		context.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": "405",
			"msg":  "Method Not Allowed",
		})
	})
	return e
}
