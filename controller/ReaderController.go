package controller

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReaderController struct {
	DB *gorm.DB
}

// QueryReader
// @Description /initreader路由
// @Author John 2023-04-16 14:55:57
// @Param ctx
func (c *ReaderController) QueryReader(ctx *gin.Context) {
	readerId := ctx.PostForm("readerId")
	if len(readerId) == 0 {
		fmt.Println("获取读者信息失败")
		response.Fail(ctx, gin.H{
			"msg":    "获取读者信息失败",
			"status": 400,
		})
		return
	}
	var reader = model.Reader{}
	if err := c.DB.Where("reader_id = ?", readerId).First(&reader).Error; err != nil {
		fmt.Println("获取读者信息失败")
		response.Fail(ctx, gin.H{
			"msg":    "获取读者信息失败",
			"status": 400,
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
	return ReaderController{
		DB: common.GetDB(),
	}
}
