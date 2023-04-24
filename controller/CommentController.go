package controller

import (
	"Go_lib/response"
	"Go_lib/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
}

// GetComments
// @Description 查询所有评论
// @Author John 2023-04-16 15:21:07
// @Param ctx
func (c *CommentController) GetComments(ctx *gin.Context) {
	commentService := service.NewCommentService()

	//c.DB.Preload("Reader").Preload("Book").Find(&comments)
	comments, lError := commentService.GetComments()
	if lError != nil {
		fmt.Println(lError.Err)
		response.Response(ctx, lError.HttpCode, gin.H{
			"status": lError.HttpCode,
			"msg":    lError.Msg,
		})
		return
	}
	response.Success(ctx, gin.H{
		"status": 100,
		"msg":    "评论区请求成功",
		"data":   comments,
	})
}

// NewCommentController
// @Description CommentController的构造器
// @Author John 2023-04-16 15:22:58
// @Return CommentController
func NewCommentController() CommentController {
	return CommentController{}
}
