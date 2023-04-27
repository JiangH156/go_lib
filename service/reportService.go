package service

import (
	"Go_lib/common"
	"Go_lib/repository"
	"Go_lib/vo"
	"errors"
	"net/http"
)

type ReportService struct {
}

// GetReports
// @Description 获取举报记录
// @Author John 2023-04-26 19:20:30
// @Param readerId
// @Return reports
// @Return lErr
func (s ReportService) GetReports(readerId string) (reports []vo.ReportVo, lErr *common.LError) {
	// 数据验证
	if readerId == "" {
		return reports, &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "获取举报记录失败",
			Err:      errors.New("数据验证失败"),
		}
	}
	//  获取举报记录
	reportRepository := repository.NewReportRepository()
	reports, err := reportRepository.GetReportsByReaderId(readerId)

	if err != nil {
		return reports, &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "获取举报记录失败",
			Err:      errors.New("获取举报记录失败"),
		}
	}
	return reports, nil
}

func NewReportService() ReportService {
	return ReportService{}
}
