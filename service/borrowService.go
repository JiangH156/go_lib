package service

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/repository"
	"Go_lib/utils"
	"Go_lib/vo"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type BorrowService struct {
	DB *gorm.DB
}

// CreateBorrowRecord
// @Description 添加借阅记录
// @Author John 2023-04-21 14:50:15
// @Param readerId
// @Param bookId
// @Param date
// @Return lErr
func (b *BorrowService) CreateBorrowRecord(readerId, bookId string, date model.Time) (lErr *common.LError) {
	tx := b.DB.Begin()
	// reader借阅次数更新
	readerRepository := repository.NewReaderRepository()
	if err := readerRepository.UpdateReaderBorrowTimes(tx, readerId, 1); err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "添加借阅记录失败",
			Err:      errors.New("reader借阅次数更新失败"),
		}
	}
	// book书籍总数更新
	bookRepository := repository.NewBookRepository()
	if err := bookRepository.UpdateBookAmount(tx, bookId, -1); err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "添加借阅记录失败",
			Err:      errors.New("book书籍总数更新失败"),
		}
	}
	// book书籍借阅次数更新
	if err := bookRepository.UpdateBookBorrowedTimes(tx, bookId, 1); err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "添加借阅记录失败",
			Err:      errors.New("book书籍借阅次数更新失败"),
		}
	}
	// reserve预约记录更新
	reserveRepository := repository.NewReserveRepository()
	updReserve := model.Reserve{
		BookId:   bookId,
		ReaderId: readerId,
		Date:     date,
	}
	if err := reserveRepository.UpdateStatus(tx, updReserve, "已借阅"); err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "添加借阅记录失败",
			Err:      errors.New("reserve预约记录更新失败"),
		}
	}
	// borrow新增记录
	addDate := time.Time(date).AddDate(0, 1, 0)
	addBorrow := model.Borrow{
		ReaderId:   readerId,
		BookId:     bookId,
		BorrowDate: date,
		ReturnDate: model.Time(addDate),
		Status:     "未还",
	}
	//fmt.Printf("%+v", addBorrow)
	borrowRepository := repository.NewBorrowRepository()
	if err := borrowRepository.CreateBorrowRecord(tx, addBorrow); err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "添加借阅记录失败",
			Err:      errors.New("borrow新增记录失败"),
		}
	}
	tx.Commit()

	return nil
}

// GetBorrows
// @Description 查询借阅记录
// @Author John 2023-04-21 23:17:19
// @Param readerId
// @Return []model.Borrow
// @Return lErr
func (b *BorrowService) GetBorrows(readerId string) (borrowVos []vo.BorrowVo, lErr *common.LError) {
	//fmt.Println(readerId)
	if readerId == "" {
		return borrowVos, &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "查询借阅记录错误",
			Err:      errors.New("readerId为空"),
		}
	}

	borrowRepository := repository.NewBorrowRepository()
	borrowVos, err := borrowRepository.GetBorrowsByReaderId(readerId)

	if err != nil {
		return borrowVos, &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "查询借阅记录错误",
			Err:      errors.New("查询借阅记录错误"),
		}
	}
	return borrowVos, nil
}

// ReturnBook
// @Description 归还书籍
// @Author John 2023-04-23 20:38:21
// @Param bookId
// @Param readerId
// @Param borrowDate
// @Return lErr
func (b *BorrowService) ReturnBook(readerId string, bookId string, borrowDate model.Time) (lErr *common.LError) {
	bookRepository := repository.NewBookRepository()
	readerRepository := repository.NewReaderRepository()
	borrowRepository := repository.NewBorrowRepository()

	if bookId == "" || readerId == "" {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "归还书籍失败",
			Err:      errors.New("bookId或readerId为空"),
		}
	}

	// 获取id
	id, err := borrowRepository.GetBorrowId(readerId, bookId, borrowDate)
	//fmt.Println(id)
	if err != nil {
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "归还书籍失败",
			Err:      errors.New("获取id失败"),
		}
	}

	// 判断书籍是否已经归还
	status, err := borrowRepository.GetBorrowStatus(id)
	if err != nil {
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "归还书籍失败",
			Err:      errors.New("查询书籍状态失败"),
		}
	}

	if status == "已还" {
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "归还书籍失败",
			Err:      errors.New("书籍已经归还"),
		}
	}

	tx := b.DB.Begin()
	// book 更新当前数量
	err = bookRepository.UpdateBookAmount(tx, bookId, 1)
	if err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "归还书籍失败",
			Err:      errors.New("book 更新当前数量失败"),
		}
	}

	// reader 判断更新逾期次数
	// 获取借阅截止时间
	returnTime, err := borrowRepository.GetBorrowReturnDate(id)
	//fmt.Println(returnTime)
	if err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "归还书籍失败",
			Err:      errors.New("获取借阅截止时间失败"),
		}
	}
	nowTime := time.Now()
	// 判断是否逾期
	if nowTime.After(time.Time(returnTime)) {
		//fmt.Println("逾期")
		// 逾期
		//status = "逾期"
		// 更新reader逾期记录
		err := readerRepository.UpdateReaderOvdTimes(tx, readerId)
		if err != nil {
			tx.Rollback()
			return &common.LError{
				HttpCode: http.StatusInternalServerError,
				Msg:      "归还书籍失败",
				Err:      errors.New("更新reader逾期记录失败"),
			}
		}
	}
	// borrow 更新实际归还日期
	realDate := model.Time(nowTime)
	err = borrowRepository.UpdateBorrowRealDate(tx, id, realDate)
	if err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "归还书籍失败",
			Err:      errors.New("borrow 更新实际归还日期失败"),
		}
	}

	// borrow 更新状态
	status = "已还"
	err = borrowRepository.UpdateBorrowStatus(tx, id, status)
	if err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "归还书籍失败",
			Err:      errors.New("borrow 更新状态失败"),
		}
	}
	tx.Commit()
	return nil
}

// RenewBook
// @Description 续借图书
// @Author John 2023-04-24 09:39:17
// @Param readerId
// @Param bookId
// @Param borrowDate
// @Return lErr
func (b *BorrowService) RenewBook(readerId string, bookId string, borrowDate string) (lErr *common.LError) {
	if readerId == "" || bookId == "" {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "数据格式有误",
			Err:      errors.New("数据格式有误"),
		}
	}
	borrowRepository := repository.NewBorrowRepository()
	t, _ := utils.ParseTime(borrowDate)
	date := model.Time(t)
	//  获取id
	id, err := borrowRepository.GetBorrowId(readerId, bookId, date)
	if err != nil {
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "续借图书失败",
			Err:      errors.New("获取id失败"),
		}
	}

	tx := b.DB.Begin()
	//更新借阅表状态-续借
	// 更新状态
	status := "续借"
	err = borrowRepository.UpdateBorrowStatus(tx, id, status)
	if err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "续借图书失败",
			Err:      errors.New("更新状态失败"),
		}
	}
	err = borrowRepository.UpdateBorrowStatus(tx, id, status)
	if err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "续借图书失败",
			Err:      errors.New("更新状态失败"),
		}
	}

	// 更新借阅时间，重置为当前时间
	broDate := model.Time(utils.NowTime())
	err = borrowRepository.UpdateBorrowBorrowDate(tx, id, broDate)
	if err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "续借图书失败",
			Err:      errors.New("更新借阅时间失败"),
		}
	}

	// 更新截止时间，重置为当前时间后一个月
	t = utils.NowTime().AddDate(0, 1, 0)
	retDate := model.Time(t)
	fmt.Println(retDate)
	err = borrowRepository.UpdateBorrowReturnDate(tx, id, retDate)
	if err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "续借图书失败",
			Err:      errors.New("更新借阅时间失败"),
		}
	}
	tx.Commit()
	return nil
}

func NewBorrowService() BorrowService {
	return BorrowService{
		DB: common.GetDB(),
	}
}
