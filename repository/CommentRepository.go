package repository

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/vo"
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

// GetCommentVos
// @Description 获取评论
// @Author John 2023-04-24 15:49:53
// @Return comments
// @Return err
func (r *CommentRepository) GetCommentVos() (comments []vo.CommentVo, err error) {
	if err = r.DB.
		Model(&model.Comment{}).
		Select(`comments.*, readers.email,readers.reader_name, books.book_name`).
		Joins("left join readers on readers.reader_id = comments.reader_id").
		Joins("left join books on books.book_id = comments.book_id").
		Scan(&comments).
		Error; err != nil {
		return comments, err
	}
	return comments, nil
}

func NewCommentRepository() CommentRepository {
	return CommentRepository{
		DB: common.GetDB(),
	}
}
