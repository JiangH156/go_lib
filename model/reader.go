package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Reader
// @Description:
type Reader struct {
	ReaderId    string `json:"readerId" gorm:"type:varchar(36);primaryKey;"`
	ReaderName  string `json:"readerName" gorm:"type:varchar(10);not null"`
	Password    string `json:"password" gorm:"type:varchar(255);not null"`
	Phone       string `json:"phone" gorm:"type:varchar(25);not null;unique"`
	BorrowTimes uint   `json:"borrowTimes"`
	OvdTimes    uint   `json:"ovdTimes"`
	Email       string `json:"email" gorm:"type:varchar(255);"`
}

// BeforeCreate
// @Description 钩子函数：插入数据前生成uuid
// @Author 2023-04-13 21:20:08
// @Param tx
// @Return err
func (r *Reader) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	r.ReaderId = uid.String()
	return nil
}
