package service

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/repository"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

// UpdateBookInfo
// @Description 管理员更新图书信息
// @Author John 2023-04-27 15:12:18
// @Param bookId
// @Param value
// @Param status
// @Param difference
// @Return lErr
func (b *BookService) UpdateBookInfo(bookId string, value string, status string, difference string) (lErr *common.LError) {
	if value == "" {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "更新图书信息失败",
			Err:      errors.New("value为空"),
		}
	}
	bookRepository := repository.NewBookRepository()
	// 开启事务
	tx := b.DB.Begin()
	switch status {
	//1 更新图书名称
	case "1":
		{
			err := bookRepository.UpdateBookNameByBookId(tx, bookId, value)
			if err != nil {
				tx.Rollback()
				return &common.LError{
					HttpCode: http.StatusInternalServerError,
					Msg:      "更新图书信息失败",
					Err:      errors.New("更新图书名称失败"),
				}
			}
		}
		//2 更新图书作者
	case "2":
		{
			err := bookRepository.UpdateAuthorByBookId(tx, bookId, value)
			if err != nil {
				tx.Rollback()
				return &common.LError{
					HttpCode: http.StatusInternalServerError,
					Msg:      "更新图书信息失败",
					Err:      errors.New("更新图书作者失败"),
				}
			}
		}
		//3 更新图书位置
	case "3":
		{
			// 判断图书位置是否使用
			//fmt.Println(value)
			book, err := bookRepository.GetBookByPosition(value)
			//fmt.Println("book====>", book)
			if book.BookId != "" {
				return &common.LError{
					HttpCode: http.StatusBadRequest,
					Msg:      "更新图书信息失败",
					Err:      errors.New("该位置已使用"),
				}
			}
			err = bookRepository.UpdatePositionByBookId(tx, bookId, value)
			if err != nil {
				tx.Rollback()
				return &common.LError{
					HttpCode: http.StatusInternalServerError,
					Msg:      "更新图书信息失败",
					Err:      errors.New("更新图书位置失败"),
				}
			}
		}
		//4 更新当前库存
	case "4":
		{
			count, err := strconv.Atoi(difference)
			if err != nil {
				return &common.LError{
					HttpCode: http.StatusBadRequest,
					Msg:      "更新图书信息失败",
					Err:      errors.New("difference转换失败"),
				}
			}
			// 更新当前数量
			err = bookRepository.UpdateAmountByBookId(tx, bookId, count)
			if err != nil {
				tx.Rollback()
				return &common.LError{
					HttpCode: http.StatusInternalServerError,
					Msg:      "更新图书信息失败",
					Err:      errors.New("更新当前数量失败"),
				}
			}
			// 更新总数量
			err = bookRepository.UpdateTotalAmountByBookId(tx, bookId, count)
			if err != nil {
				tx.Rollback()
				return &common.LError{
					HttpCode: http.StatusInternalServerError,
					Msg:      "更新图书信息失败",
					Err:      errors.New("更新总数量失败"),
				}
			}
		}
	}
	tx.Commit()
	return nil
}

// DeleteBook
// @Description 管理员删除书籍
// @Author John 2023-04-27 20:39:08
// @Param bookId
func (b *BookService) DeleteBook(bookId string) (lErr *common.LError) {
	// 数据验证
	if bookId == "" {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "删除书籍失败",
			Err:      errors.New("数据验证失败"),
		}
	}

	bookRepository := repository.NewBookRepository()

	// 查询书籍当前库存
	amount, err := bookRepository.GetAmountByBookId(bookId)
	if err != nil {
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "删除书籍失败",
			Err:      errors.New("查询当前库存失败"),
		}
	}
	// 查询书籍总库存
	totalAmount, err := bookRepository.GetTotalAmountByBookId(bookId)
	if err != nil {
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "删除书籍失败",
			Err:      errors.New("查询书籍总库存失败"),
		}
	}
	//fmt.Println(amount, totalAmount)
	// 比较当前库存和总库存是否相等
	if amount != totalAmount {
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "当前书籍存在未归还书籍",
			Err:      errors.New("当前书籍存在未归还书籍"),
		}
	}
	// 开启事务
	tx := b.DB.Begin()
	err = bookRepository.DeleteBookByBookId(tx, bookId)
	if err != nil {
		// 事务回滚
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "删除书籍失败",
			Err:      errors.New("删除书籍失败"),
		}
	}
	return nil
}

func NewBookService() BookService {
	return BookService{
		DB: common.GetDB(),
	}
}
