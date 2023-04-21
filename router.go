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
	r.POST("/books", bookController.GetBooks)
	r.POST("/searchbook", bookController.GetBookInfo)

	// comment
	commentController := controller.NewCommentController()
	r.POST("/comments", commentController.GetComments)

	// reader
	readerController := controller.NewReaderController()
	r.POST("/initreader", readerController.GetReaderInfo)

	// borrow
	borrowController := controller.NewBorrowController()
	r.POST("/addborrow", borrowController.CreateBorrowRecord)

	// reserve
	reserveController := controller.NewReserveController()
	r.POST("/addreserve", reserveController.CreateReserveRecord)
	r.POST("/reserve", reserveController.GetReserves)
	r.POST("/cancelreserve", reserveController.DeleteReserveRecord)

	return r
}
