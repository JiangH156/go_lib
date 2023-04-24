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

// UpdateReaderBorrowTimes
// @Description 更新借阅次数
// @Author John 2023-04-21 15:40:41
// @Param tx
// @Param readerId
func (r *ReaderRepository) UpdateReaderBorrowTimes(tx *gorm.DB, readerId string, count int) error {
	return tx.Model(&model.Reader{}).Where("reader_id = ?", readerId).UpdateColumn("borrow_times", gorm.Expr("borrow_times + ?", count)).Error
}

// UpdateReaderOvdTimes
// @Description 更新逾期记录
// @Author John 2023-04-23 22:01:09
// @Param tx
// @Param readerId
func (r *ReaderRepository) UpdateReaderOvdTimes(tx *gorm.DB, readerId string) error {
	if err := tx.
		Model(&model.Reader{}).
		Where("reader_id = ?", readerId).
		UpdateColumn("ovd_times", gorm.Expr("ovd_times + ?", 1)).
		Error; err != nil {
		return err
	}
	return nil
}

func NewReaderRepository() ReaderRepository {
	return ReaderRepository{
		DB: common.GetDB(),
	}
}
