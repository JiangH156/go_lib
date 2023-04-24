package service

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/repository"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type BookService struct {
	DB *gorm.DB
}

// GetBooks
// @Description 查询所有书籍
// @Author John 2023-04-20 20:51:45
// @Return []model.Book
// @Return *common.LError
func (b *BookService) GetBooks() (books []model.Book, lErr *common.LError) {
	bookRepository := repository.NewBookRepository()
	books, err := bookRepository.GetBooks()
	if err != nil {
		return books, &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "书籍查询失败",
			Err:      err,
		}
	}
	// 请求书籍数据为空
	if len(books) == 0 {
		return books, &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "请求书籍数据为空",
			Err:      errors.New("请求书籍数据为空"),
		}
	}
	return books, nil
}

// GetBookByName
// @Description 查询书籍
// @Author John 2023-04-20 20:51:57
func (b *BookService) GetBookByName(bookName string) (books []model.Book, lErr *common.LError) {
	var bookRepository = repository.NewBookRepository()
	books, err := bookRepository.GetBooksByName(bookName)
	// 查询出错
	if err != nil {
		return books, &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "查询书籍错误",
			Err:      err,
		}
	}
	return books, nil
}

func NewBookService() BookService {
	return BookService{
		DB: common.GetDB(),
	}
}
