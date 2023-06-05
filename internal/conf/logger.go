package conf

import (
	"io"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func newFileLogger() io.Writer {
	return &lumberjack.Logger{
		Filename:  loggerFileSetting.SavePath + "/" + loggerFileSetting.FileName + loggerFileSetting.FileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}
}

func setupLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(loggerSetting.logLevel())

	out := newFileLogger()
	logrus.SetOutput(out)
}
