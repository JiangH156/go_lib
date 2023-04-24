package controller

import (
	"Go_lib/response"
	"Go_lib/service"
	"Go_lib/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type BorrowController struct {
}

// CreateBorrowRecord
// @Description 新增借阅记录
// @Author John 2023-04-20 16:01:38
func (b *BorrowController) CreateBorrowRecord(ctx *gin.Context) {
	borrowService := service.NewBorrowService()

	readerId := ctx.PostForm("readerId")
	bookId := ctx.PostForm("bookId")
	date := ctx.PostForm("date")

	addTime, _ := utils.ParseTime(date)

	lErr := borrowService.CreateBorrowRecord(readerId, bookId, addTime)
	if lErr != nil {
		fmt.Println(lErr.Err)
		response.Response(ctx, lErr.HttpCode, gin.H{
			"status": lErr.HttpCode,
			"msg":    lErr.Msg,
		})
		return
	}
	response.Success(ctx, gin.H{
		"status": 200,
		"msg":    "新增借阅记录",
	})
}

// GetBorrows
// @Description 获取所有借阅记录
// @Author John 2023-04-21 23:12:12
// @Param ctx
func (b *BorrowController) GetBorrows(ctx *gin.Context) {
	readerId := ctx.PostForm("readerId")
	borrowService := service.NewBorrowService()

	borrows, lErr := borrowService.GetBorrows(readerId)

	if lErr != nil {
		fmt.Println(lErr.Err)
		response.Response(ctx, lErr.HttpCode, gin.H{
			"status": lErr.HttpCode,
			"msg":    lErr.Msg,
		})
		return
	}
	response.Success(ctx, gin.H{
		"status": 200,
		"msg":    "获取借阅记录成功",
		"data":   borrows,
	})
}

// ReturnBook
// @Description 归还书籍
// @Author John 2023-04-23 20:13:18
// @Param ctx
func (b *BorrowController) ReturnBook(ctx *gin.Context) {
	borrowService := service.NewBorrowService()
	bookId := ctx.PostForm("bookId")
	readerId := ctx.PostForm("readerId")
	borrowDate := ctx.PostForm("borrowDate")
	date, _ := utils.ParseTime(borrowDate)
	lErr := borrowService.ReturnBook(readerId, bookId, date)
	if lErr != nil {
		fmt.Println(lErr.Err)
		response.Response(ctx, lErr.HttpCode, gin.H{
			"status": lErr.HttpCode,
			"msg":    lErr.Msg,
		})
		return
	}

	response.Success(ctx, gin.H{
		"status": 200,
		"msg":    "归还书籍成功",
	})

}

// RenewBook
// @Description 续借图书
// @Author John 2023-04-24 09:31:20
// @Param ctx
func (b *BorrowController) RenewBook(ctx *gin.Context) {
	readerId := ctx.PostForm("readerId")
	bookId := ctx.PostForm("bookId")
	borrowDate := ctx.PostForm("borrowDate")

	borrowService := service.NewBorrowService()
	lErr := borrowService.RenewBook(readerId, bookId, borrowDate)
	if lErr != nil {
		fmt.Println(lErr.Err)
		response.Response(ctx, lErr.HttpCode, gin.H{
			"status": lErr.HttpCode,
			"msg":    lErr.Msg,
		})
		return
	}
	response.Success(ctx, gin.H{
		"status": 200,
		"msg":    "续借图书成功",
	})
}

// NewBorrowController
// @Description BookController的构造器
// @Author John 2023-04-16 15:23:25
// @Return BookController
func NewBorrowController() BorrowController {
	return BorrowController{}
}
