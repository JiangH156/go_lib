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

	// book
	bookController := controller.NewBookController()
	r.POST("/books", bookController.QueryBooks)
	r.POST("/searchbook", bookController.QueryBook)

	// comment
	commentController := controller.NewCommentController()
	r.POST("/comments", commentController.QueryComments)

	// reader
	readerController := controller.NewReaderController()
	r.POST("/initreader", readerController.QueryReader)

	// borrow
	borrowController := controller.NewBorrowController()
	r.POST("/addborrow", borrowController.AddBorrow)

	// reserve
	reserveController := controller.NewReserveController()
	r.POST("/addreserve", reserveController.AddReserve)
	r.POST("/reserve", reserveController.QueryReserveByReaderId)
	r.POST("/cancelreserve", reserveController.DeleteReserve)

	return r
}
