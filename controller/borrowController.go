package controller

import (
	"Go_lib/model"
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

	t, _ := utils.ParseTime(date)
	addTime := model.Time(t)
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

// GetReaderBorrowRecords
// @Description 获取所有借阅记录
// @Author John 2023-04-21 23:12:12
// @Param ctx
func (b *BorrowController) GetReaderBorrowRecords(ctx *gin.Context) {
	readerId := ctx.PostForm("readerId")
	borrowService := service.NewBorrowService()
	borrowVos, lErr := borrowService.GetReaderBorrowRecords(readerId)

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
		"data":   borrowVos,
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

	d, _ := utils.ParseTime(borrowDate)
	date := model.Time(d)
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

// GetAllBorrowRecords
// @Description 获取全部借阅记录
// @Author John 2023-04-27 21:32:39
// @Param ctx
func (b *BorrowController) GetAllBorrowRecords(ctx *gin.Context) {
	borrowService := service.NewBorrowService()
	borrowVos, lErr := borrowService.GetAllBorrowRecords()

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
		"data":   borrowVos,
	})
}

// GetBorrowRecordByInfo
// @Description 管理员 根据关键词获取相关借阅记录
// @Author John 2023-04-27 22:49:00
// @Param ctx
func (b *BorrowController) GetBorrowRecordByInfo(ctx *gin.Context) {
	info := ctx.PostForm("info")
	borrowService := service.NewBorrowService()
	borrowVos, lErr := borrowService.GetBorrowRecordByInfo(info)

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
		"data":   borrowVos,
	})
}

// DeleteBorrow
// @Description 管理员删除借阅记录
// @Author John 2023-04-27 22:58:00
// @Param ctx
func (b *BorrowController) DeleteBorrow(ctx *gin.Context) {
	readerId := ctx.PostForm("readerId")
	bookId := ctx.PostForm("bookId")
	borrowDate := ctx.PostForm("borrowDate")

	borrowService := service.NewBorrowService()
	lErr := borrowService.DeleteBorrow(readerId, bookId, borrowDate)

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
	})
}

// SendReminder
// @Description 管理员提醒用户还书
// @Author John 2023-04-28 09:45:29
// @Param ctx
func (b *BorrowController) SendReminder(ctx *gin.Context) {
	readerId := ctx.PostForm("readerId")
	bookName := ctx.PostForm("bookName")

	borrowService := service.NewBorrowService()
	lErr := borrowService.SendReminder(readerId, bookName)
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
	})
}

func NewBorrowController() BorrowController {
	return BorrowController{}
}
