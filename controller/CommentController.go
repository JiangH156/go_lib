package controller

import (
	"Go_lib/common"
	"Go_lib/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
	DB *gorm.DB
}

// QueryComments
// @Description 查询所有评论
// @Author John 2023-04-16 15:21:07
// @Param ctx
func (c *CommentController) QueryComments(ctx *gin.Context) {
	var comments = []model.Comment{}

	//c.DB.Preload("Reader").Preload("Book").Find(&comments)
	c.DB.
		Select("readers.email, comments.status, comment_id,comments.reader_id,comments.book_id,comments.reader_id, readers.reader_name,books.book_name, date, content, praise").
		Joins("left join readers on readers.reader_id = comments.reader_id").
		Joins("left join books on books.book_id = comments.reader_id").
		Find(&comments)
	ctx.JSON(200, gin.H{
		"msg":    "评论区请求成功",
		"status": 200,
		"data":   comments,
	})
}

// NewCommentController
// @Description CommentController的构造器
// @Author John 2023-04-16 15:22:58
// @Return CommentController
func NewCommentController() CommentController {
	return CommentController{
		DB: common.GetDB(),
	}
}
