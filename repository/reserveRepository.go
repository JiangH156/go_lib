package repository

import (
	"fmt"
	"github.com/jiangh156/Go_lib/common"
	"github.com/jiangh156/Go_lib/model"
	"github.com/jiangh156/Go_lib/vo"
	"gorm.io/gorm"
)

type ReserveRepository struct {
	DB *gorm.DB
}

// CreateReserveRecord
// @Description 添加预约记录
// @Author John 2023-04-20 14:40:35
// @Param addReserve
// @Return error
func (r *ReserveRepository) CreateReserveRecord(tx *gorm.DB, addReserve model.Reserve) error {
	if err := tx.Create(&addReserve).Error; err != nil {
		fmt.Println("添加预约记录失败")
		// 出错回滚
		return err
	}
	return nil
}

// GetReserveVosByReaderId
// @Description 根据readerID查询ReserveVO
// @Author John 2023-04-20 22:06:08
// @Param readerId
// @Return []vo.ReserveVO
// @Return error
func (r *ReserveRepository) GetReserveVosByReaderId(readerId string) (reserveVos []vo.ReserveVo, err error) {
	err = r.DB.
		Table("reserves").
		Select(`reserves.reader_id, books.book_id, reserves.status, books.author, books.book_name, reserves.date`).
		Joins(`JOIN books ON reserves.book_id = books.book_id`).
		Where(`reserves.reader_id = ?`, readerId).
		Scan(&reserveVos).
		Error
	if err != nil {
		return nil, err
	}
	return reserveVos, nil
}

// GetReserveById
// @Description  根据id查询Reserve
// @Author John 2023-04-20 22:10:51
// @Param readerId
// @Param bookId
// @Return reserve
// @Return err
func (r *ReserveRepository) GetReserveById(id string) (reserve model.Reserve, err error) {
	err = r.DB.Where("id = ?", id).First(&reserve).Error
	if err != nil {
		return reserve, err
	}
	return reserve, nil
}

// DeleteReserveRecordById
// @Description 删除预约记录
// @Author John 2023-04-20 14:43:27
// @Param bookId
// @Param readerId
// @Param date
// @Return error
func (r *ReserveRepository) DeleteReserveRecordById(tx *gorm.DB, id string) error {
	if err := tx.Where("id = ?", id).
		Delete(&model.Reserve{}).
		Error; err != nil {
		return err
	}
	return nil
}

// UpdateStatus
// @Description 更新预约记录状态
// @Author John 2023-04-21 20:14:50
// @Param tx
// @Param reserve
// @Param status
// @Return error
func (r *ReserveRepository) UpdateStatus(tx *gorm.DB, reserve model.Reserve, status string) error {
	return tx.
		Model(model.Reserve{}).
		Where("reader_id = ?", reserve.ReaderId).
		Where("book_id = ?", reserve.BookId).
		Where("date = ?", reserve.Date).
		UpdateColumn("status", status).
		Error
}

// GetReserveId
// @Description 获取id
// @Author John 2023-04-24 16:09:58
// @Param readerId
// @Param bookId
// @Param date
// @Return id
// @Return err
func (r *ReserveRepository) GetReserveId(readerId string, bookId string, date model.Time) (id string, err error) {
	if err = r.DB.
		Model(&model.Reserve{}).
		Select("id").
		Where("reader_id = ?", readerId).
		Where("book_id = ?", bookId).
		Where("date = ?", date).
		Scan(&id).
		Error; err != nil {
		return id, err
	}
	return id, nil
}

// GetAllReserveRecords
// @Description 管理员获取所有预约记录
// @Author John 2023-04-28 14:55:12
// @Return reserveVos
// @Return lErr
func (r *ReserveRepository) GetAllReserveRecords() (reserveVos []vo.ReserveVo, err error) {
	if err = r.DB.
		Table("reserves").
		Select(`readers.reader_name, books.book_name, reserves.date, reserves.reader_id, reserves.book_id`).
		Joins(`JOIN readers ON readers.reader_id = reserves.reader_id`).
		Joins(`JOIN books ON books.book_id = reserves.book_id`).
		Scan(&reserveVos).
		Error; err != nil {
		return reserveVos, err
	}
	return reserveVos, nil
}

func NewReserveRepository() ReserveRepository {
	return ReserveRepository{
		DB: common.GetDB(),
	}
}
