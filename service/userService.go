package service

import (
	"Go_lib/common"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService() UserService {
	return UserService{
		DB: common.GetDB(),
	}
}
