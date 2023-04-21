package service

import (
	"Go_lib/common"
	"gorm.io/gorm"
)

type BorrowService struct {
	DB *gorm.DB
}

func NewBorrowService() BorrowService {
	return BorrowService{
		DB: common.GetDB(),
	}
}
