package repository

import (
	"github.com/jiangh156/Go_lib/common"
	"github.com/jiangh156/Go_lib/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// GetReaderByPhone
// @Description 根据电话号码获取读者
// @Author John 2023-05-05 14:58:36
// @Param phone
// @Return reader
// @Return err
func (r *UserRepository) GetReaderByPhone(phone string) (reader model.Reader, err error) {
	if err = r.DB.Where("phone = ?", phone).First(&reader).Error; err != nil {
		return reader, err
	}
	return reader, nil
}

// CreateReader
// @Description 新增读者
// @Author John 2023-05-05 15:06:31
// @Param tx
// @Param reader
// @Return error
func (r *UserRepository) CreateReader(tx *gorm.DB, reader model.Reader) error {
	if err := tx.Create(&reader).Error; err != nil {
		return err
	}
	return nil
}

// GetAdminByPhone
// @Description 通过手机号获取管理员
// @Author John 2023-05-05 15:20:04
// @Param phone
// @Return model.Admin
// @Return bool
func (r *UserRepository) GetAdminByPhone(phone string) (admin model.Admin, exist bool) {
	r.DB.Where("phone = ?", phone).First(&admin)
	return admin, admin.Id != 0
}

func NewUserRepository() UserRepository {
	return UserRepository{
		DB: common.GetDB(),
	}
}
