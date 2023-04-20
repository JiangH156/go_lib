package controller

import (
	"Go_lib/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BorrowController struct {
	DB *gorm.DB
}

// AddBorrow
// @Description 新增借阅记录
// @Author John 2023-04-20 16:01:38
func (b *BorrowController) AddBorrow(ctx *gin.Context) {
	
}

// NewBorrowController
// @Description BookController的构造器
// @Author John 2023-04-16 15:23:25
// @Return BookController
func NewBorrowController() BorrowController {
	return BorrowController{
		DB: common.GetDB(),
	}
}
