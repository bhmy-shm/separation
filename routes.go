package main

import (
	"Gin/前后端分离/controller"
	"Gin/前后端分离/middleware"

	"github.com/gin-gonic/gin"
)

//
func CollectRoute(r *gin.Engine) *gin.Engine {

	//注册
	r.POST("/api/auth/register", controller.Register)

	//登录
	r.POST("/api/auth/login", controller.Login)

	//创建用户信息的路由
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
