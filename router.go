package main

import (
	"Go_lib/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// 配置CORS跨域路由
	CORSMiddleware := middleware.CORSMiddleware()
	r.Use(CORSMiddleware)

	return r
}
