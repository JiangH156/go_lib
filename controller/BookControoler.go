package controller

import (
	"Go_lib/response"
	"Go_lib/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookController struct {
}

// GetBooks
// @Description 查询所有书籍
// @Author John 2023-04-15 15:36:55
// @Param ctx
func (b *BookController) GetBooks(ctx *gin.Context) {
	bookService := service.NewBookService()
	books, lErr := bookService.GetBooks()
	// 查询错误
	if lErr != nil {
		fmt.Println(lErr.Err)
		response.Response(ctx, lErr.HttpCode, gin.H{
			"status": lErr.HttpCode,
			"msg":    lErr.Msg,
		})
		return
	}
	response.Response(ctx, http.StatusOK, gin.H{
		"status": 200,
		"msg":    "书籍请求成功",
		"data":   books,
	})
}

// GetBooksByName
// @Description 查询书籍
// @Author John 2023-04-18 15:33:55
// @Param ctx
func (b *BookController) GetBooksByName(ctx *gin.Context) {
	bookService := service.NewBookService()
	name := ctx.PostForm("name")
	// name为空，跳转到QueryBooks
	if name == "" {
		ctx.Redirect(http.StatusFound, "/books")
		ctx.Abort()
	}
	books, lErr := bookService.GetBookByName(name)
	// 查询出错
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
		"msg":    "查询成功",
		"data":   books,
	})
}

// NewBookController
// @Description  BookController的构造器
// @Author John 2023-04-16 15:21:28
// @Return BookController
func NewBookController() BookController {
	return BookController{}
}
