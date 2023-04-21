package service

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/repository"
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

func NewReaderService() ReaderService {
	return ReaderService{
		DB: common.GetDB(),
	}
}
