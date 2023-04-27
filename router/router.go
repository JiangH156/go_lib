package router

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
	r.POST("/searchbook", bookController.GetBooksByName)
	r.POST("/changebookinfo", bookController.UpdateBookInfo)
	r.POST("/delbook", bookController.DeleteBook)

	// comment
	commentController := controller.NewCommentController()
	r.POST("/comments", commentController.GetComments)
	r.POST("/amount", commentController.GetCommentCount)
	r.POST("/addcomment", commentController.CreateComment)

	// reader
	readerController := controller.NewReaderController()
	r.POST("/initreader", readerController.GetReaderInfo)
	r.POST("/amountmax", readerController.MaxCountReader)
	r.POST("/delperson", readerController.DeleteReader)

	// admin
	r.POST("/initreaderlist", readerController.GetReaders)

	// borrow
	borrowController := controller.NewBorrowController()
	r.POST("/addborrow", borrowController.CreateBorrowRecord)
	r.POST("/borrows", borrowController.GetBorrows)
	r.POST("/returnbook", borrowController.ReturnBook)
	r.POST("/continueborrow", borrowController.RenewBook)

	// reserve
	reserveController := controller.NewReserveController()
	r.POST("/addreserve", reserveController.CreateReserveRecord)
	r.POST("/reserve", reserveController.GetReserves)
	r.POST("/cancelreserve", reserveController.DeleteReserveRecord)

	// report
	reportController := controller.NewReportController()
	r.POST("/initstureport", reportController.GetReports)

	return r
}
