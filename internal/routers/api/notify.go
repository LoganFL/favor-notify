package api

import (
	"favor-notify/internal/services"
	"favor-notify/pkg/app"
	"favor-notify/pkg/errcode"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func PushNotify(c *gin.Context) {
	param := services.ReqPushNotify{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	err := services.PushNotify(param)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(nil)
}

func PushNotifyDao(c *gin.Context) {
	param := services.ReqPushNotify{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	err := services.PushNotifyDao(param)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(nil)
}

func PushNotifySys(c *gin.Context) {
	param := services.ReqPushNotifySys{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	err := services.PushNotifySys(param)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(nil)
}
