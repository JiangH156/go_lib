package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jiangh156/Go_lib/response"
	"github.com/jiangh156/Go_lib/service"
)

type ReaderController struct {
}

// GetReaderInfo
// @Description 查询读者信息
// @Author John 2023-04-16 14:55:57
// @Param ctx
func (c *ReaderController) GetReaderInfo(ctx *gin.Context) {
	readerId := ctx.PostForm("readerId")
	readerService := service.NewReaderService()
	// 查询数据
	reader, lErr := readerService.GetReader(readerId)

	//判断查询是否出错
	if lErr != nil {
		fmt.Println(lErr.Err)
		response.Response(ctx, lErr.HttpCode, gin.H{
			"status": lErr.HttpCode,
			"msg":    lErr.Msg,
		})
		return
	}

	response.Success(ctx, gin.H{
		"status":      200,
		"msg":         "获取读者信息成功",
		"readerId":    reader.ReaderId,
		"readerName":  reader.ReaderName,
		"readerPhone": reader.Phone,
		"borrowTimes": reader.BorrowTimes,
		"ovdTimes":    reader.OvdTimes,
		"email":       reader.Email,
		"isAdmin":     false,
	})
}

// GetMaxCountReader
// @Description 获取评论最多的人员
// @Author John 2023-04-25 21:29:09
// @Param ctx
func (c *ReaderController) GetMaxCountReader(ctx *gin.Context) {
	readerService := service.NewReaderService()
	reader, lErr := readerService.MaxCountReader()

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
		"msg":    "请求成功",
		"data":   reader,
	})

}

// GetReaders
// @Description 管理员获取所有人员信息
// @Author John 2023-04-26 21:54:20
// @Param ctx
func (c *ReaderController) GetReaders(ctx *gin.Context) {
	readerService := service.NewReaderService()
	readers, lErr := readerService.GetReaders()

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
		"msg":    "请求成功",
		"data":   readers,
	})
}

// DeleteReader
// @Description 管理员删除用户
// @Author John 2023-04-26 22:18:42
// @Param ctx
func (c *ReaderController) DeleteReader(ctx *gin.Context) {
	// 数据接收
	readerId := ctx.PostForm("readerId")

	readerService := service.NewReaderService()
	lErr := readerService.DeleteReader(readerId)
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
		"msg":    "请求成功",
	})
}

func NewReaderController() ReaderController {
	return ReaderController{}
}
