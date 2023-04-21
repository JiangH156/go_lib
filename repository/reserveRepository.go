package repository

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/vo"
	"fmt"
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
func (r *ReserveRepository) CreateReserveRecord(addReserve model.Reserve) error {
	// 开启事务
	tx := r.DB.Begin()
	if err := tx.Create(&addReserve).Error; err != nil {
		fmt.Println("添加预约记录失败")
		// 出错回滚
		tx.Rollback()
		return err
	}
	// 数据插入成功,提交事务
	tx.Commit()
	return nil
}

// GetReserveVOsByReaderId
// @Description 根据readerID查询ReserveVO
// @Author John 2023-04-20 22:06:08
// @Param readerId
// @Return []vo.ReserveVO
// @Return error
func (r *ReserveRepository) GetReserveVOsByReaderId(readerId string) (reserveVOs []vo.ReserveVO, err error) {
	err = r.DB.
		Table("reserves").
		Select(`reserves.reader_id, books.book_id, reserves.status, books.author, books.book_name, reserves.date`).
		Joins(`JOIN books ON reserves.book_id = books.book_id`).
		Where(`reserves.reader_id = ?`, readerId).
		Scan(&reserveVOs).
		Error
	if err != nil {
		return nil, err
	}
	return reserveVOs, nil
}

// GetReserveByReaderIDAndBookID
// @Description  根据readerID和bookId查询Reserve
// @Author John 2023-04-20 22:10:51
// @Param readerId
// @Param bookId
// @Return reserve
// @Return err
func (r *ReserveRepository) GetReserveByReaderIDAndBookID(readerId, bookId string) (reserve model.Reserve, err error) {
	err = r.DB.Where("reader_id = ?", readerId).Where("book_id = ?", bookId).First(&reserve).Error
	if err != nil {
		return reserve, err
	}
	return reserve, nil
}

// DeleteReserveRecord
// @Description 删除预约记录
// @Author John 2023-04-20 14:43:27
// @Param bookId
// @Param readerId
// @Param date
// @Return error
func (r *ReserveRepository) DeleteReserveRecord(bookId string, readerId string, date model.Time) error {
	//开启事务
	tx := r.DB.Begin()
	if err := tx.Where("book_id = ?", bookId).
		Where("reader_id = ?", readerId).
		Where("date = ?", date).
		Delete(&model.Reserve{}).
		Error; err != nil {
		//删除失败
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func NewReserveRepository() ReserveRepository {
	return ReserveRepository{
		DB: common.GetDB(),
	}
}
