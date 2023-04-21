package repository

import (
	"Go_lib/common"
	"Go_lib/model"
	"gorm.io/gorm"
)

type ReaderRepository struct {
	DB *gorm.DB
}

// GetReaderByReaderId
// @Description 根据ReaderId查询读者信息
// @Author John 2023-04-20 21:30:10
// @Param readerId
// @Return model.Reader
// @Return error
func (r *ReaderRepository) GetReaderByReaderId(readerId string) (*model.Reader, error) {
	var reader = model.Reader{}
	err := r.DB.Where("reader_id = ?", readerId).First(&reader).Error
	return &reader, err
}

func NewReaderRepository() ReaderRepository {
	return ReaderRepository{
		DB: common.GetDB(),
	}
}
