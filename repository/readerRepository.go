package repository

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/vo"
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

// MaxCountReader
// @Description 返回评论最多的读者
// @Author John 2023-04-25 21:34:28
// @Return comment
// @Return err
func (r *ReaderRepository) MaxCountReader() (reader []vo.MaxCountReader, err error) {
	if err = r.DB.
		Table("readers").
		Select(`COUNT(*) as amount, reader_name`).
		Joins("LEFT JOIN comments ON readers.reader_id = comments.reader_id").
		Group("readers.reader_id").
		Order("amount DESC").
		Limit(1).
		Scan(&reader).
		Error; err != nil {
		return reader, err
	}
	return reader, nil
}

// GetReaders
// @Description 返回所有人员信息
// @Author John 2023-04-26 22:12:22
// @Return readers
// @Return lErr
func (r *ReaderRepository) GetReaders() (readers []model.Reader, err error) {
	if err = r.DB.Model(&model.Reader{}).Find(&readers).Error; err != nil {
		return readers, err
	}
	return readers, nil
}

// DeleteReaderByReaderId
// @Description 管理员删除读者
// @Author John 2023-04-26 22:22:08
// @Param readerId
// @Return err
func (r *ReaderRepository) DeleteReaderByReaderId(tx *gorm.DB, readerId string) (err error) {
	if err = tx.Where("reader_id = ?", readerId).Delete(&model.Reader{}).Error; err != nil {
		return err
	}
	return nil
}

func NewReaderRepository() ReaderRepository {
	return ReaderRepository{
		DB: common.GetDB(),
	}
}
