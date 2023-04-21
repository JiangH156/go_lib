package repository

import (
	"Go_lib/common"
	"Go_lib/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

// GetBooksByName
// @Description 数据层，根据书名查询书籍
// @Author John 2023-04-18 16:41:27
// @Param name
// @Return []model.Book
// @Return error
func (b *BookRepository) GetBooksByName(bookName string) ([]model.Book, error) {
	var books = []model.Book{}
	if err := b.DB.Where("book_name like ?", "%"+bookName+"%").Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func NewBookRepository() BookRepository {
	return BookRepository{
		DB: common.GetDB(),
	}
}
