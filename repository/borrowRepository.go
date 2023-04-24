package repository

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/vo"
	"gorm.io/gorm"
)

type BorrowRepository struct {
	DB *gorm.DB
}

// CreateBorrowRecord
// @Description 创建借阅记录
// @Author John 2023-04-21 20:45:45
// @Param borrow
// @Return error
func (r *BorrowRepository) CreateBorrowRecord(tx *gorm.DB, borrow model.Borrow) error {
	return tx.Create(&borrow).Error
}

// GetBorrowsByReaderId
// @Description 通过readerId获取借阅记录
// @Author John 2023-04-21 23:20:19
// @Param id
func (r *BorrowRepository) GetBorrowsByReaderId(readerId string) (borrowVos []vo.BorrowVo, err error) {
	if err := r.DB.
		Table("borrows").
		Select(`borrows.*, books.book_name, books.author`).
		Joins(`left join books on books.book_id = borrows.book_id`).
		Scan(&borrowVos).
		Error; err != nil {
		return borrowVos, err
	}

	return borrowVos, nil
}

// GetBorrowReturnDate
// @Description 返回借阅截止时间
// @Author John 2023-04-23 21:15:49
// @Param readerId
// @Param bookId
// @Param borrowDate
// @Return error
func (r *BorrowRepository) GetBorrowReturnDate(id string) (returnTime model.Time, err error) {
	if err := r.DB.
		Model(&model.Borrow{}).
		Select(`return_date`).
		Where("id = ?", id).
		Scan(&returnTime).
		Error; err != nil {
		return returnTime, err
	}
	return returnTime, nil
}

// GetBorrowId
// @Description 返回Id
// @Author John 2023-04-23 22:24:48
// @Param readerId
// @Param bookId
// @Param borrowDate
// @Return id
// @Return err
func (r *BorrowRepository) GetBorrowId(readerId string, bookId string, borrowDate model.Time) (id string, err error) {
	if err := r.DB.
		Model(&model.Borrow{}).
		Select(`id`).
		Where("reader_id = ?", readerId).
		Where("book_id = ?", bookId).
		Where("borrow_date = ?", borrowDate).
		Scan(&id).
		Error; err != nil {
		return id, err
	}
	return id, nil
}

// UpdateBorrowRealDate
// @Description 更新实际归还时间
// @Author John 2023-04-23 22:29:10
// @Param tx
// @Param id
// @Param date
// @Return error
func (r *BorrowRepository) UpdateBorrowRealDate(tx *gorm.DB, id string, realDate model.Time) error {
	if err := tx.
		Model(&model.Borrow{}).
		Where("id = ?", id).
		UpdateColumn("real_date", realDate).
		Error; err != nil {
		return err
	}
	return nil
}

// UpdateBorrowStatus
// @Description 更新状态
// @Author John 2023-04-23 22:32:43
// @Param tx
// @Param id
// @Param status
// @Return error
func (r *BorrowRepository) UpdateBorrowStatus(tx *gorm.DB, id string, status string) error {
	if err := tx.
		Model(&model.Borrow{}).
		Where("id = ?", id).
		UpdateColumn("status", status).
		Error; err != nil {
		return err
	}
	return nil
}

// GetBorrowStatus
// @Description 返回状态
// @Author John 2023-04-23 23:18:59
// @Param id
// @Return status
// @Return err
func (r *BorrowRepository) GetBorrowStatus(id string) (status string, err error) {
	if err = r.DB.
		Model(&model.Borrow{}).
		Select(`status`).
		Where("id = ?", id).
		Scan(&status).Error; err != nil {
		return status, err
	}
	return status, err
}

func NewBorrowRepository() BorrowRepository {
	return BorrowRepository{
		DB: common.GetDB(),
	}
}
