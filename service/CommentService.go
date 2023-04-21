package service

import (
	"Go_lib/common"
	"Go_lib/vo"
	"gorm.io/gorm"
)

type CommentService struct {
	DB *gorm.DB
}

// GetComments
// @Description 查询所有评论
// @Author John 2023-04-20 21:10:09
// @Return []vo.CommentVO
// @Return *common.LError
func (c *CommentService) GetComments() (comments []vo.CommentVO, lErr *common.LError) {
	//c.DB.Preload("Reader").Preload("Book").Find(&comments)
	c.DB.
		Select("readers.email, comments.status, comment_id,comments.reader_id,comments.book_id,comments.reader_id, readers.reader_name,books.book_name, date, content, praise").
		Joins("left join readers on readers.reader_id = comments.reader_id").
		Joins("left join books on books.book_id = comments.reader_id").
		Find(&comments)
	return nil, nil
}

func NewCommentService() CommentService {
	return CommentService{
		DB: common.GetDB(),
	}
}
