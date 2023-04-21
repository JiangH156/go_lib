package controller

import (
	"Go_lib/response"
	"Go_lib/service"
	"fmt"
	"github.com/gin-gonic/gin"
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

// NewReaderController
// @Description ReaderController的构造器
// @Author John 2023-04-16 15:22:48
// @Return ReaderController
func NewReaderController() ReaderController {
	return ReaderController{}
}
