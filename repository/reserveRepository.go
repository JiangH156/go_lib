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

// AddReserve
// @Description 添加预约记录
// @Author John 2023-04-20 14:40:35
// @Param addReserve
// @Return error
func (r *ReserveRepository) AddReserve(addReserve model.Reserve) error {
	// 开启事务
	tx := r.DB.Begin()
	if err := tx.Create(&addReserve).Error; err != nil {
		fmt.Println("预约失败")
		// 出错回滚
		tx.Rollback()
		return err
	}
	// 数据插入成功,提交事务
	tx.Commit()
	return nil
}

// QueryReserveByReaderId
// @Description 查询预约记录
// @Author John 2023-04-20 14:40:30
// @Param readerId
// @Return []vo.ReserveVO
// @Return error
func (r *ReserveRepository) QueryReserveByReaderId(readerId string) ([]vo.ReserveVO, error) {
	var reserveVO = []vo.ReserveVO{}
	err := r.DB.
		Table("reserves").
		Select(`reserves.reader_id, books.book_id, reserves.status, books.author, books.book_name, reserves.date`).
		Joins(`JOIN books ON reserves.book_id = books.book_id`).
		Where(`reserves.reader_id = ?`, readerId).
		Scan(&reserveVO).
		Error
	if err != nil {
		return nil, err
	}
	return reserveVO, nil
}

// DeleteReserve
// @Description 删除预约记录
// @Author John 2023-04-20 14:43:27
// @Param bookId
// @Param readerId
// @Param date
// @Return error
func (r *ReserveRepository) DeleteReserve(bookId string, readerId string, date model.Time) error {
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
