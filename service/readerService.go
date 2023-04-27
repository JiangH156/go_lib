package service

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/repository"
	"Go_lib/vo"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type ReaderService struct {
	DB *gorm.DB
}

// GetReader
// @Description 查询读者信息
// @Author John 2023-04-20 21:14:23
// @Param ctx
func (c *ReaderService) GetReader(readerId string) (reader *model.Reader, lErr *common.LError) {
	readerRepository := repository.NewReaderRepository()
	if len(readerId) == 0 {
		return reader, &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "获取读者信息失败",
			Err:      errors.New("获取读者信息失败"),
		}
	}
	// 查询reader
	reader, err := readerRepository.GetReaderByReaderId(readerId)
	// 查询错误
	if err != nil {
		return reader, &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "获取读者信息失败",
			Err:      errors.New("获取读者信息失败"),
		}
	}
	return reader, nil
}

// MaxCountReader
// @Description 获取评论最多的人员
// @Author John 2023-04-25 20:58:11
func (c *ReaderService) MaxCountReader() (reader []vo.MaxCountReader, lErr *common.LError) {
	readerRepository := repository.NewReaderRepository()
	reader, err := readerRepository.MaxCountReader()

	//  查询评论数量
	if err != nil {
		return reader, &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "请求失败",
			Err:      errors.New("查询评论数量失败"),
		}
	}
	return reader, nil
}

// GetReaders
// @Description 返回所有人员信息
// @Author John 2023-04-26 22:11:36
// @Return readers
// @Return lErr
func (c *ReaderService) GetReaders() (readers []model.Reader, lErr *common.LError) {
	readerRepository := repository.NewReaderRepository()
	// 获取所有人员信息
	readers, err := readerRepository.GetReaders()

	if err != nil {
		return readers, &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "请求失败",
			Err:      errors.New("获取所有人员信息失败"),
		}
	}
	return readers, nil
}

// DeleteReader
// @Description 管理员删除用户
// @Author John 2023-04-26 22:21:12
// @Param readerId
// @Return lErr
func (c *ReaderService) DeleteReader(readerId string) (lErr *common.LError) {
	// 获取未归还的书籍
	borrowRepository := repository.NewBorrowRepository()
	borrows, err := borrowRepository.GetUnreturnedBorrowsByReaderId(readerId)
	if err != nil {
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "请求失败",
			Err:      errors.New("获取未归还的书籍失败"),
		}
	}
	// 判断未归还书籍数量
	//fmt.Println("borrow", borrows)
	if len(borrows) != 0 {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "请求失败",
			Err:      errors.New("用户存在未归还书籍"),
		}
	}

	tx := c.DB.Begin()
	readerRepository := repository.NewReaderRepository()
	err = readerRepository.DeleteReaderByReaderId(tx, readerId)
	if err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "请求失败",
			Err:      errors.New("删除用户失败"),
		}
	}
	//tx.Commit()
	return nil
}

func NewReaderService() ReaderService {
	return ReaderService{
		DB: common.GetDB(),
	}
}
