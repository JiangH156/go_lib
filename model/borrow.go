package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Borrow
// @Description:
type Borrow struct {
	Id       string `json:"id" gorm:"type:varchar(36);primaryKey"`
	ReaderId string `json:"readerId" gorm:"type:varchar(36);primaryKey;"`
	// Reader外键
	//Reader `gorm:"foreignKey:ReaderId;references:ReaderId"`
	BookId string `json:"bookId" gorm:"type:varchar(50);primaryKey"`
	// Book外键
	//Book `gorm:"foreignKey:BookId;references:BookId"`
	// 借阅日期
	BorrowDate Time `json:"borrowDate"`
	// 截止日期
	ReturnDate Time `json:"ReturnDate"`
	// 实际归还日期
	RealDate Time `json:"RealDate"`
	// 借阅状态:借出、已还、续借
	Status string `json:"status" gorm:"type:varchar(255)"`
}

func (r *Borrow) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	r.Id = uid.String()
	return nil
}
