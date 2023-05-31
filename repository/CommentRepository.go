package repository

import (
	"github.com/jiangh156/Go_lib/common"
	"github.com/jiangh156/Go_lib/model"
	"github.com/jiangh156/Go_lib/vo"
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

// GetCommentCount
// @Description 返回评论数量
// @Author John 2023-04-25 20:06:56
// @Return count
// @Return err
func (r *CommentRepository) GetCommentCount() (count int64, err error) {
	if err = r.DB.Model(&model.Comment{}).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

// CreateComment
// @Description 添加评论
// @Author John 2023-04-25 22:19:46
// @Param comment
// @Return error
func (r *CommentRepository) CreateComment(tx *gorm.DB, comment model.Comment) (err error) {
	if err := tx.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

// GetCommentId
// @Description 返回commnetId
// @Author John 2023-04-28 16:21:15
// @Param readerId
// @Param bookId
// @Param date
// @Return commentId
// @Return err
func (r *CommentRepository) GetCommentId(readerId string, bookId string, date model.Time) (commentId string, err error) {
	if err = r.DB.
		Model(&model.Comment{}).
		Select(`comment_id`).
		Where(`reader_id = ? AND book_id = ? AND date = ?`, readerId, bookId, date).
		Scan(&commentId).
		Error; err != nil {
		return commentId, err
	}
	return commentId, nil
}

// UpdatePraiseByCommentId
// @Description 更新点赞记录
// @Author John 2023-04-28 16:27:29
// @Param tx
// @Param commentId
// @Return err
func (r *CommentRepository) UpdatePraiseByCommentId(tx *gorm.DB, commentId string) (err error) {
	if err = tx.
		Model(&model.Comment{}).
		Where("comment_id = ?", commentId).
		UpdateColumn("praise", gorm.Expr("praise + ?", 1)).
		Error; err != nil {
		return err
	}
	return nil
}

// GetReaderIdByCommentId
// @Description 获取书名ID
// @Author John 2023-05-03 14:55:13
// @Param commentId
// @Return readerId
func (r *CommentRepository) GetReaderIdByCommentId(commentId string) (readerId string, err error) {
	if err = r.DB.
		Model(&model.Comment{}).
		Select("reader_id").
		Where("comment_id = ?", commentId).
		Scan(&readerId).
		Error; err != nil {
		return readerId, err
	}
	return readerId, nil
}

// UpdateStatusByCommentId
// @Description 更新Status字段
// @Author John 2023-05-03 21:36:16
// @Param tx
// @Param commentId
// @Return error
func (r *CommentRepository) UpdateStatusByCommentId(tx *gorm.DB, commentId string, status int) error {
	if err := tx.
		Model(&model.Comment{}).
		Where("comment_id = ?", commentId).
		UpdateColumn("status", status).
		Error; err != nil {
		return err
	}
	return nil
}

func NewCommentRepository() CommentRepository {
	return CommentRepository{
		DB: common.GetDB(),
	}
}
