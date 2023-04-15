package main

import (
	"Go_lib/controller"
	"Go_lib/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// 配置CORS跨域路由
	CORSMiddleware := middleware.CORSMiddleware()
	r.Use(CORSMiddleware)

	// user
	userController := controller.NewUserController()
	r.POST("/login", userController.Login)
	r.POST("/register", userController.Register)

	return r
}
