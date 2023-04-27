package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	// 该评论信息的唯一标识
	CommentId string `json:"commentId" gorm:"type:varchar(50);primary_key"`
	// 评论人的唯一标识
	ReaderId string `json:"readerId" gorm:"type:varchar(36)"`
	// 外键关联
	//Reader `gorm:"foreignKey:ReaderId;references:ReaderId"`
	// 被评论书籍的唯一标识
	BookId string `json:"bookId" gorm:"type:varchar(50)"`
	// 外键关联
	//Book `gorm:"foreignKey:BookId;references:BookId"`
	// 评论日期
	Date Time `json:"date"`
	// 评论内容
	Content string `json:"content" gorm:"type:varchar(255);"`
	// 点赞数
	Praise uint `json:"praise"`
	// 评论状态
	// 1：正常  3： 删帖
	Status uint `json:"status"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	c.CommentId = uid.String()
	return nil
}
