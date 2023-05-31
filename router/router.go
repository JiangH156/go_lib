package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jiangh156/Go_lib/controller"
	"github.com/jiangh156/Go_lib/middleware"
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
	r.POST("/adminaddbooks", bookController.CreateBook)

	// comment
	commentController := controller.NewCommentController()
	r.POST("/comments", commentController.GetComments)
	r.POST("/amount", commentController.GetCommentCount)
	r.POST("/addcomment", commentController.CreateComment)
	r.POST("/addpraise", commentController.UpdatePraise)

	// reader
	readerController := controller.NewReaderController()
	r.POST("/initreader", readerController.GetReaderInfo)
	r.POST("/amountmax", readerController.GetMaxCountReader)
	r.POST("/delperson", readerController.DeleteReader)
	r.POST("/initreaderlist", readerController.GetReaders)

	// borrow
	borrowController := controller.NewBorrowController()
	r.POST("/addborrow", borrowController.CreateBorrowRecord)
	r.POST("/borrows", borrowController.GetReaderBorrowRecords)
	r.POST("/returnbook", borrowController.ReturnBook)
	r.POST("/continueborrow", borrowController.RenewBook)
	r.POST("/borrowslist", borrowController.GetAllBorrowRecords)
	r.POST("/searchborrow", borrowController.GetBorrowRecordByInfo)
	r.POST("/deleteborrow", borrowController.DeleteBorrow)
	r.POST("/alertperson", borrowController.SendReminder)

	// reserve
	reserveController := controller.NewReserveController()
	r.POST("/addreserve", reserveController.CreateReserveRecord)
	r.POST("/reserve", reserveController.GetReserveRecords)
	r.POST("/cancelreserve", reserveController.DeleteReserveRecord)
	r.POST("/reservelist", reserveController.GetAllReserveRecords)

	// report
	reportController := controller.NewReportController()
	r.POST("/initstureport", reportController.GetReportRecords)
	r.POST("/initreportlist", reportController.GetAllReportRecords)
	r.POST("/reportcomment", reportController.CreateReport)
	r.POST("/auditcomment", reportController.ManageReport)

	return r
}
