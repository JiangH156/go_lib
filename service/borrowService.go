package service

import (
	"errors"
	"fmt"
	"github.com/John/Go_lib/common"
	"github.com/John/Go_lib/model"
	"github.com/John/Go_lib/repository"
	"github.com/John/Go_lib/utils"
	"github.com/John/Go_lib/vo"
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

// GetReaderBorrowRecords
// @Description 查询借阅记录
// @Author John 2023-04-21 23:17:19
// @Param readerId
// @Return []model.Borrow
// @Return lErr
func (b *BorrowService) GetReaderBorrowRecords(readerId string) (borrowVos []vo.BorrowVo, lErr *common.LError) {
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

// GetAllBorrowRecords
// @Description 获取全部借阅记录
// @Author John 2023-04-27 22:11:53
// @Return borrowVos
// @Return lErr
func (b *BorrowService) GetAllBorrowRecords() (borrowVos []vo.BorrowVo, lErr *common.LError) {
	borrowRepository := repository.NewBorrowRepository()
	borrowVos, err := borrowRepository.GetAllBorrowRecords()

	if err != nil {
		return borrowVos, &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "请求失败",
			Err:      errors.New("获取全部借阅记录错误"),
		}
	}
	return borrowVos, nil
}

// GetBorrowRecordByInfo
// @Description 根据关键词获取相关借阅记录
// @Author John 2023-04-27 22:50:08
// @Param info
// @Return borrowVos
// @Return lErr
func (b *BorrowService) GetBorrowRecordByInfo(info string) (borrowVos []vo.BorrowVo, lErr *common.LError) {
	borrowRepository := repository.NewBorrowRepository()
	borrowVos, err := borrowRepository.GetBorrowRecordByInfo(info)

	if err != nil {
		return borrowVos, &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "请求失败",
			Err:      errors.New("获取全部借阅记录错误"),
		}
	}
	return borrowVos, nil
}

// DeleteBorrow
// @Description 管理员删除借阅记录
// @Author John 2023-04-27 23:06:11
// @Param readerId
// @Param bookId
// @Param borrowDate
// @Return borrowVos
// @Return lErr
func (b *BorrowService) DeleteBorrow(readerId, bookId, borrowDate string) (lErr *common.LError) {
	// 数据验证
	if readerId == "" || bookId == "" || borrowDate == "" {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "请求失败",
			Err:      errors.New("删除借阅记录错误"),
		}
	}
	date, _ := utils.ParseTime(borrowDate)
	delBorrow := model.Borrow{
		ReaderId:   readerId,
		BookId:     bookId,
		BorrowDate: model.Time(date),
	}

	// 开启事务
	tx := b.DB.Begin()
	borrowRepository := repository.NewBorrowRepository()
	err := borrowRepository.DeleteBorrow(tx, delBorrow)

	if err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "请求失败",
			Err:      errors.New("删除借阅记录错误"),
		}
	}
	tx.Commit()
	return nil
}

// SendReminder
// @Description 管理员提醒用户还书
// @Author John 2023-04-28 09:47:44
// @Param readerId
// @Param bookName
// @Return lErr
func (b *BorrowService) SendReminder(readerId string, bookName string) (lErr *common.LError) {
	// 数据验证
	if readerId == "" || bookName == "" {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "数据验证失败",
			Err:      errors.New("数据验证失败"),
		}
	}

	borrowRepository := repository.NewBorrowRepository()
	// 查询用户未归还书籍
	borrows, err := borrowRepository.GetUnreturnedBorrowsByReaderId(readerId)
	if err != nil {
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "请求失败",
			Err:      errors.New("查询用户未归还书籍错误"),
		}
	}
	//fmt.Println(borrows)
	//  用户是否归还书籍
	if len(borrows) == 0 {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "用户已归还书籍",
			Err:      errors.New("用户已归还书籍"),
		}
	}
	// 获取书籍id
	bookRepository := repository.NewBookRepository()
	bookId, err := bookRepository.GetBookIdByBookName(bookName)
	flag := false
	for _, b := range borrows {
		if b.BookId == bookId {
			// 用户未归还书籍中存在该书籍
			flag = true
		}
	}
	//  书籍已归还
	if !flag {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "用户已归还书籍",
			Err:      errors.New("用户已归还书籍"),
		}
	}

	// 获取邮箱
	readerRepository := repository.NewReaderRepository()
	reader, err := readerRepository.GetReaderByReaderId(readerId)
	if err != nil {
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "请求失败",
			Err:      errors.New("获取邮箱失败"),
		}
	}
	// 发送邮件
	subject := "还书提醒！"
	body := "读者您好，请尽快归还书籍:" + bookName
	err = utils.SendEmail([]string{reader.Email}, nil, nil, subject, body, "")
	if err != nil {
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "请求失败",
			Err:      errors.New("发送邮件失败"),
		}
	}

	return nil
}

func NewBorrowService() BorrowService {
	return BorrowService{
		DB: common.GetDB(),
	}
}
