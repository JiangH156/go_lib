package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Reserve
// @Description: 预约记录
type Reserve struct {
	Id       string `json:"ID" gorm:"type:varchar(36);primaryKey;"`
	ReaderId string `json:"readerId" gorm:"type:varchar(50);"`
	BookId   string `json:"bookId" gorm:"type:varchar(50);"`
	Date     Time   `json:"date"`
	Status   string `json:"status" gorm:"type:varchar(50)"`
}

// BeforeCreate
// @Description 插入数据前添加uuid
// @Author John 2023-04-16 16:57:07
// @Param tx
// @Return err
func (r *Reserve) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	r.Id = uid.String()
	return nil
}
