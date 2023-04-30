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

// UpdateBookNameByBookId
// @Description 更新书名
// @Author John 2023-04-27 15:39:07
// @Param tx
// @Param bookId
// @Param bookName
// @Return error
func (b *BookRepository) UpdateBookNameByBookId(tx *gorm.DB, bookId string, bookName string) error {
	if err := tx.Model(&model.Book{}).Where("book_id = ?", bookId).UpdateColumn("book_name", bookName).Error; err != nil {
		return err
	}
	return nil
}

// UpdateAuthorByBookId
// @Description 更新作者
// @Author John 2023-04-27 15:40:10
// @Param tx
// @Param bookId
// @Param author
// @Return interface{}
func (b *BookRepository) UpdateAuthorByBookId(tx *gorm.DB, bookId string, author string) interface{} {
	if err := tx.Model(&model.Book{}).Where("book_id = ?", bookId).UpdateColumn("author", author).Error; err != nil {
		return err
	}
	return nil
}

// UpdatePositionByBookId
// @Description  更新书籍位置
// @Author John 2023-04-27 15:45:40
// @Param tx
// @Param bookId
// @Param position
// @Return error
func (b *BookRepository) UpdatePositionByBookId(tx *gorm.DB, bookId string, position string) error {
	if err := tx.Model(&model.Book{}).Where("book_id = ?", bookId).UpdateColumn("position", position).Error; err != nil {
		return err
	}
	return nil
}

// GetBookByPosition
// @Description 返回指定位置的图书
// @Author John 2023-04-27 15:48:38
// @Param position
// @Return book
// @Return err
func (b *BookRepository) GetBookByPosition(position string) (book model.Book, err error) {
	if err = b.DB.Model(&model.Book{}).Where("position = ?", position).First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

// UpdateTotalAmountByBookId
// @Description 更新总数量
// @Author John 2023-04-27 15:58:14
// @Param tx
// @Param bookId
// @Param count
// @Return error
func (b *BookRepository) UpdateTotalAmountByBookId(tx *gorm.DB, bookId string, count int) error {
	if err := tx.
		Model(&model.Book{}).
		Where("book_id = ?", bookId).
		UpdateColumn("amount", gorm.Expr("amount + ?", count)).
		Error; err != nil {
		return err
	}
	return nil
}

// UpdateAmountByBookId
// @Description 更新当前数量
// @Author John 2023-04-27 15:58:17
// @Param tx
// @Param bookId
// @Param count
// @Return error
func (b *BookRepository) UpdateAmountByBookId(tx *gorm.DB, bookId string, count int) error {
	if err := tx.
		Model(&model.Book{}).
		Where("book_id = ?", bookId).
		UpdateColumn("total_amount", gorm.Expr("total_amount + ?", count)).
		Error; err != nil {
		return err
	}
	return nil
}

// DeleteBookByBookId
// @Description 根据书籍id删除书籍
// @Author John 2023-04-27 20:42:04
// @Param tx
// @Param bookId
// @Return error
func (b *BookRepository) DeleteBookByBookId(tx *gorm.DB, bookId string) error {
	if err := tx.Where("book_id = ?", bookId).Delete(&model.Book{}).Error; err != nil {
		return err
	}
	return nil
}

// GetAmountByBookId
// @Description 返回当前书籍当前库存
// @Author John 2023-04-27 20:46:36
// @Param bookId
// @Return amount
// @Return err
func (b *BookRepository) GetAmountByBookId(bookId string) (amount int, err error) {
	if err = b.DB.Model(&model.Book{}).Select(`amount`).Where("book_id = ?", bookId).Scan(&amount).Error; err != nil {
		return amount, err
	}
	return amount, nil
}

// GetTotalAmountByBookId
// @Description  返回当前书籍总库存
// @Author John 2023-04-27 20:48:27
// @Param bookId
// @Return totalAmount
// @Return err
func (b *BookRepository) GetTotalAmountByBookId(bookId string) (totalAmount int, err error) {
	if err = b.DB.Model(&model.Book{}).Select(`total_amount`).Where("book_id = ?", bookId).Scan(&totalAmount).Error; err != nil {
		return totalAmount, err
	}
	return totalAmount, nil
}

// GetBookIdByBookName
// @Description 根据书籍名称获取书籍id
// @Author John 2023-04-30 10:33:54
// @Param bookName
// @Return interface{}
func (b *BookRepository) GetBookIdByBookName(bookName string) (bookId string, err error) {
	if err = b.DB.Model(&model.Book{}).Select(`book_id`).Where("book_name = ?", bookName).Scan(&bookId).Error; err != nil {
		return bookId, err
	}
	return bookId, nil
}

func NewBookRepository() BookRepository {
	return BookRepository{
		DB: common.GetDB(),
	}
}
