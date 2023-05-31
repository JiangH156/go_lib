package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jiangh156/Go_lib/response"
	"github.com/jiangh156/Go_lib/service"
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
		"status": 200,
		"msg":    "评论区请求成功",
		"data":   comments,
	})
}

// GetCommentCount
// @Description 获取评论数量
// @Author John 2023-04-25 20:02:05
// @Param ctx
func (c *CommentController) GetCommentCount(ctx *gin.Context) {
	commentService := service.NewCommentService()
	mytotal, lErr := commentService.GetCommentAmount()

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
		"data":   mytotal,
	})
}

// CreateComment
// @Description 添加评论
// @Author John 2023-04-25 22:01:06
// @Param ctx
func (c *CommentController) CreateComment(ctx *gin.Context) {
	// 接收数据
	readerId := ctx.PostForm("readerId")
	bookId := ctx.PostForm("bookId")
	content := ctx.PostForm("content")
	commentService := service.NewCommentService()
	lErr := commentService.CreateComment(readerId, bookId, content)
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

// UpdatePraise
// @Description 更新点赞记录
// @Author John 2023-04-28 16:12:31
// @Param ctx
func (c *CommentController) UpdatePraise(ctx *gin.Context) {
	// 接收数据
	readerId := ctx.PostForm("readerId")
	bookId := ctx.PostForm("bookId")
	date := ctx.PostForm("date")
	commentService := service.NewCommentService()
	lErr := commentService.UpdatePraise(readerId, bookId, date)
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

// NewCommentController
// @Description CommentController的构造器
// @Author John 2023-04-16 15:22:58
// @Return CommentController
func NewCommentController() CommentController {
	return CommentController{}
}
