package main

import (
	"Gin/前后端分离/controller"

	"github.com/gin-gonic/gin"
)

//
func CollectRoute(r *gin.Engine) *gin.Engine {

	r.POST("/api/auth/register", controller.Register)

	return r
}
