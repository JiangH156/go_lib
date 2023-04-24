package repository

import (
	"Go_lib/common"
	"Go_lib/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

// GetBooks
// @Description 查询所有书籍
// @Author John 2023-04-21 21:09:03
// @Return books
// @Return err
func (b *BookRepository) GetBooks() (books []model.Book, err error) {
	if err := b.DB.Find(&books).Error; err != nil {
		return books, err
	}
	//fmt.Println(books)
	return books, nil
}

// GetBooksByName
// @Description 根据书名查询书籍
// @Author John 2023-04-18 16:41:27
// @Param name
// @Return []model.Book
// @Return error
func (b *BookRepository) GetBooksByName(bookName string) (books []model.Book, err error) {
	if err := b.DB.Where("book_name like ?", "%"+bookName+"%").Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

// UpdateBookAmount
// @Description 更新书籍总数
// @Author John 2023-04-21 16:28:17
// @Param tx
// @Param bookId
// @Param count
// @Return error
func (b *BookRepository) UpdateBookAmount(tx *gorm.DB, bookId string, count int) error {
	return tx.Model(&model.Book{}).Where("book_id = ?", bookId).UpdateColumn("amount", gorm.Expr("amount + ?", count)).Error
}

// UpdateBookBorrowedTimes
// @Description 更新书籍借阅次数
// @Author John 2023-04-21 16:28:39
// @Param tx
// @Param id
// @Param i
// @Return error
func (b *BookRepository) UpdateBookBorrowedTimes(tx *gorm.DB, bookId string, count int) error {
	return tx.Model(&model.Book{}).Where("book_id = ?", bookId).UpdateColumn("borrowed_times", gorm.Expr("borrowed_times + ?", count)).Error
}

func NewBookRepository() BookRepository {
	return BookRepository{
		DB: common.GetDB(),
	}
}
