package service

import (
	"Go_lib/common"
	"Go_lib/repository"
	"Go_lib/vo"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type CommentService struct {
	DB *gorm.DB
}

// GetComments
// @Description 查询所有评论
// @Author John 2023-04-20 21:10:09
// @Return []vo.CommentVO
// @Return *common.LError
func (c *CommentService) GetComments() (comments []vo.CommentVo, lErr *common.LError) {
	//c.DB.Preload("Reader").Preload("Book").Find(&comments)
	commentRepository := repository.NewCommentRepository()
	comments, err := commentRepository.GetCommentVos()
	// 获取评论CommentVo
	if err != nil {
		return comments, &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "查询所有评论失败",
			Err:      errors.New("获取评论CommentVo失败"),
		}
	}
	return comments, nil
}

func NewCommentService() CommentService {
	return CommentService{
		DB: common.GetDB(),
	}
}
