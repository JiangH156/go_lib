package controller

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/repository"
	"Go_lib/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type BookController struct {
	DB *gorm.DB
}

// QueryBooks
// @Description 查询所有书籍 /books
// @Author John 2023-04-15 15:36:55
// @Param ctx
func (b *BookController) QueryBooks(ctx *gin.Context) {
	books := []model.Book{}
	if err := b.DB.Find(&books).Error; err != nil {
		fmt.Println("书籍请求失败")
		response.Response(ctx, http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "书籍请求失败",
		})
		return
	}
	// 请求书籍数据为空
	if len(books) == 0 {
		fmt.Println("请求书籍数据为空")
		response.Response(ctx, http.StatusBadRequest, gin.H{
			"status": 400,
			"msg":    "请求书籍数据为空",
		})
		return
	}
	response.Response(ctx, http.StatusOK, gin.H{
		"status": 200,
		"msg":    "书籍请求成功",
		"data":   books,
	})
}

// QueryBook
// @Description 根据关键词搜索书籍 /searchbook
// @Author John 2023-04-18 15:33:55
// @Param ctx
func (b *BookController) QueryBook(ctx *gin.Context) {
	var books = []model.Book{}
	var bookRepository = repository.NewBookRepository()
	name := ctx.PostForm("name")
	// name为空，跳转到QueryBooks
	if name == "" {
		ctx.Redirect(http.StatusFound, "/books")
		ctx.Abort()
	}
	books, err := bookRepository.QueryBooksByName(name)
	// 查询出错
	if err != nil {
		fmt.Println("查询书籍信息错误")
		response.Response(ctx, http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "查询书籍信息错误",
		})
		return
	}
	// 查询数据判断
	if len(books) == 0 {
		fmt.Println("查询结果为空")
		response.Response(ctx, http.StatusBadRequest, gin.H{
			"status": 400,
			"msg":    "查询结果为空",
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
	return BookController{
		DB: common.GetDB(),
	}
}
