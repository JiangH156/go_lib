package service

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/repository"
	"Go_lib/vo"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

type ReserveService struct {
	DB *gorm.DB
}

// CreateReserveRecord
// @Description 新增预约记录
// @Author John 2023-04-20 22:05:10
// @Param reserve
// @Return lErr
func (r *ReserveService) CreateReserveRecord(addReserve model.Reserve) (lErr *common.LError) {
	reserveRepository := repository.NewReserveRepository()
	if addReserve.ReaderId == "" || addReserve.BookId == "" {
		fmt.Println("预约失败")
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "预约失败",
			Err:      errors.New("预约失败"),
		}
	}

	// 验证数据库是否已经存在该预约
	reserve, _ := reserveRepository.GetReserveByReaderIDAndBookID(addReserve.ReaderId, addReserve.BookId)
	if reserve.Id != "" {
		fmt.Println("预约记录已存在")
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "预约记录已存在",
			Err:      errors.New("预约记录已存在"),
		}
	}

	if err := reserveRepository.CreateReserveRecord(addReserve); err != nil {
		fmt.Println(err)
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "新增预约记录失败",
			Err:      errors.New("新增预约记录失败"),
		}
	}
	return nil
}

// GetReserves
// @Description 根据readerId获取预约信息
// @Author John 2023-04-20 22:52:29
// @Param readerId
// @Return reserveVOs
// @Return lErr
func (r *ReserveService) GetReserves(readerId string) (reserveVOs []vo.ReserveVO, lErr *common.LError) {
	var reserveRepository = repository.NewReserveRepository()
	if readerId == "" {
		return reserveVOs, &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "查询预约记录失败",
			Err:      errors.New("readerId为空"),
		}
	}

	reserveVOs, err := reserveRepository.GetReserveVOsByReaderId(readerId)
	if err != nil {
		return reserveVOs, &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "读者请求预约记录失败",
			Err:      errors.New("读者请求预约记录失败"),
		}
	}

	//查询数据为空
	if len(reserveVOs) == 0 {
		fmt.Println("读者请求预约记录为空")
		return reserveVOs, &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "读者请求预约记录为空",
			Err:      errors.New("读者请求预约记录为空"),
		}
	}
	return reserveVOs, nil
}

// DeleteReserveRecord
// @Description 删除预约记录
// @Author John 2023-04-20 22:59:06
// @Param delReserve
// @Return lErr
func (r *ReserveService) DeleteReserveRecord(delReserve model.Reserve) (lErr *common.LError) {
	reserveRepository := repository.NewReserveRepository()
	if err := reserveRepository.DeleteReserveRecord(delReserve.BookId, delReserve.ReaderId, delReserve.Date); err != nil {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "取消预约失败",
			Err:      errors.New("取消预约失败"),
		}
	}
	return nil
}

func NewReserveService() ReserveService {
	return ReserveService{
		DB: common.GetDB(),
	}
}
