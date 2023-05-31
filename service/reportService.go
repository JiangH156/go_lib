package service

import (
	"errors"
	"github.com/John/Go_lib/common"
	"github.com/John/Go_lib/model"
	"github.com/John/Go_lib/repository"
	"github.com/John/Go_lib/utils"
	"github.com/John/Go_lib/vo"
	"gorm.io/gorm"
	"net/http"
)

type ReportService struct {
	DB *gorm.DB
}

// GetReportRecords
// @Description 获取举报记录
// @Author John 2023-04-26 19:20:30
// @Param readerId
// @Return reports
// @Return lErr
func (r *ReportService) GetReportRecords(readerId string) (reports []vo.ReportVo, lErr *common.LError) {
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
	reports, err := reportRepository.GetReportRecordsByReaderId(readerId)

	if err != nil {
		return reports, &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "获取举报记录失败",
			Err:      errors.New("获取举报记录失败"),
		}
	}
	return reports, nil
}

// GetAllReportRecords
// @Description 获取所有举报记录
// @Author John 2023-04-28 15:12:12
// @Return reports
// @Return lErr
func (r *ReportService) GetAllReportRecords() (reportVos []vo.ReportVo, lErr *common.LError) {
	//  获取举报记录
	reportRepository := repository.NewReportRepository()
	reportVos, err := reportRepository.GetAllReportRecords()

	if err != nil {
		return reportVos, &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "获取举报记录失败",
			Err:      errors.New("获取举报记录失败"),
		}
	}
	return reportVos, nil
}

// CreateReport
// @Description 用户举报评论
// @Author John 2023-05-02 11:43:58
// @Param commentId
// @Param reporterId
// @Return lErr
func (r *ReportService) CreateReport(commentId string, reporterId string) (lErr *common.LError) {
	// 数据验证
	if commentId == "" || reporterId == "" {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "数据验证失败",
			Err:      errors.New("数据验证失败"),
		}
	}

	// 判断举报人和被举报评论是否属于同一个人
	commentRepository := repository.NewCommentRepository()
	readerId, err := commentRepository.GetReaderIdByCommentId(commentId)
	if err != nil {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "请求失败",
			Err:      errors.New("查询ReaderId失败"),
		}
	}
	//  不允许举报自己
	if readerId == reporterId {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "不允许举报自己",
			Err:      errors.New("不允许举报自己"),
		}
	}

	// 不允许举报管理员
	if readerId == "admin" {
		return &common.LError{
			HttpCode: http.StatusBadRequest,
			Msg:      "不允许举报管理员",
			Err:      errors.New("不允许举报管理员"),
		}
	}

	// 开启事务
	tx := r.DB.Begin()
	reportRepository := repository.NewReportRepository()
	t := utils.NowTime()
	creReport := model.Report{
		CommentId:  commentId,
		ReporterId: reporterId,
		ReportDate: model.Time(t),
		Status:     "审核中",
	}
	err = reportRepository.CreateReport(tx, creReport)
	if err != nil {
		tx.Rollback()
		return &common.LError{
			HttpCode: http.StatusInternalServerError,
			Msg:      "请求失败",
			Err:      errors.New("数据库插入数据错误"),
		}
	}
	tx.Commit()
	return nil
}

// ManageReport
// @Description 管理评论举报信息
// @Author John 2023-05-03 09:58:52
// @Param reportInfo
// @Return lErr
func (c *ReportService) ManageReport(reportInfo vo.ReportVo) (lErr *common.LError) {

	commentRepository := repository.NewCommentRepository()
	reportRepository := repository.NewReportRepository()
	readerRepository := repository.NewReaderRepository()

	// 开启事务
	tx := c.DB.Begin()

	//管理员删除举报评论
	if reportInfo.Status == "0" {
		//数据验证
		if utils.IsAnyParameterEmpty(reportInfo.CommentID, reportInfo.ReporterID,
			reportInfo.ReportDate, reportInfo.ReaderID, reportInfo.BookID, reportInfo.Date) {
			return &common.LError{
				HttpCode: http.StatusBadRequest,
				Msg:      "请求参数有误",
				Err:      errors.New("请求参数有误"),
			}
		}

		// 获取CommentId
		commentId, err := commentRepository.GetCommentId(reportInfo.ReaderID, reportInfo.BookID, reportInfo.Date)
		if err != nil {
			return &common.LError{
				HttpCode: http.StatusInternalServerError,
				Msg:      "请求失败",
				Err:      errors.New("获取CommentId错误"),
			}
		}
		// 更新Comment
		err = commentRepository.UpdateStatusByCommentId(tx, commentId, 0)
		if err != nil {
			tx.Rollback()
			return &common.LError{
				HttpCode: http.StatusInternalServerError,
				Msg:      "请求失败",
				Err:      errors.New("更新Comment错误"),
			}
		}
		// 更新Report
		report := model.Report{
			CommentId:  reportInfo.CommentID,
			ReporterId: reportInfo.ReporterID,
			ReportDate: reportInfo.ReportDate,
			Status:     "已通过",
		}
		err = reportRepository.UpdateStatus(tx, report)
		if err != nil {
			tx.Rollback()
			return &common.LError{
				HttpCode: http.StatusInternalServerError,
				Msg:      "请求失败",
				Err:      errors.New("更新Report错误"),
			}
		}
		// 举报人反馈
		email, err := readerRepository.GetEmailByReaderId(reportInfo.ReporterID)
		if err != nil {
			//tx.Rollback()
			return &common.LError{
				HttpCode: http.StatusInternalServerError,
				Msg:      "请求失败",
				Err:      errors.New("获取邮箱错误"),
			}
		}
		subject := "举报成功！"
		body := "我们已经对该用户的不良行为进行处理，感谢您对社区做出的贡献！"
		err = utils.SendEmail([]string{email}, nil, nil, subject, body, "")
		if err != nil {
			tx.Rollback()
			return &common.LError{
				HttpCode: http.StatusInternalServerError,
				Msg:      "请求失败",
				Err:      err,
			}
		}

		// 被举报人反馈
		email, err = readerRepository.GetEmailByReaderId(reportInfo.ReaderID)
		if err != nil {
			tx.Rollback()
			return &common.LError{
				HttpCode: http.StatusInternalServerError,
				Msg:      "请求失败",
				Err:      errors.New("获取邮箱错误"),
			}
		}
		subject = "警告！！"
		body = "我们收到用户对您的举报，希望您能遵守秩序，文明用语！"
		err = utils.SendEmail([]string{email}, nil, nil, subject, body, "")
		if err != nil {
			tx.Rollback()
			return &common.LError{
				HttpCode: http.StatusInternalServerError,
				Msg:      "请求失败",
				Err:      err,
			}
		}
	}

	//管理员驳回举报评论
	if reportInfo.Status == "1" {
		// 数据验证
		if utils.IsAnyParameterEmpty(reportInfo.CommentID, reportInfo.ReporterID) {
			return &common.LError{
				HttpCode: http.StatusInternalServerError,
				Msg:      "请求参数有误",
				Err:      errors.New("请求参数有误"),
			}
		}
		report := model.Report{
			CommentId:  reportInfo.CommentID,
			ReporterId: reportInfo.ReporterID,
			ReportDate: reportInfo.ReportDate,
			Status:     "已驳回",
		}
		// 驳回该举报
		err := reportRepository.UpdateStatus(tx, report)
		if err != nil {
			tx.Rollback()
			return &common.LError{
				HttpCode: http.StatusInternalServerError,
				Msg:      "请求失败",
				Err:      errors.New("驳回该举报错误"),
			}
		}
		// 获取邮箱(ReaderId 为reporterId）
		email, err := readerRepository.GetEmailByReaderId(reportInfo.ReporterID)
		if err != nil || email == "" {
			tx.Rollback()
			return &common.LError{
				HttpCode: http.StatusInternalServerError,
				Msg:      "请求失败",
				Err:      errors.New("获取邮箱错误"),
			}
		}
		//发送反馈邮件
		subject := "举报反馈！"
		body := "我们暂无检测到该用户的不良行为，感谢您为保护社区环境做出的贡献！"
		err = utils.SendEmail([]string{email}, nil, nil, subject, body, "")
		if err != nil {
			tx.Rollback()
			return &common.LError{
				HttpCode: http.StatusInternalServerError,
				Msg:      "请求失败",
				Err:      err,
			}
		}
	}

	//管理员自行评论
	if reportInfo.Status == "3" {
		// 数据验证
		if utils.IsAnyParameterEmpty(reportInfo.ReaderID, reportInfo.BookID, reportInfo.Date) {
			return &common.LError{
				HttpCode: http.StatusBadRequest,
				Msg:      "请求参数有误",
				Err:      errors.New("请求参数有误"),
			}
		}
		// 获取CommentId
		commentId, err := commentRepository.GetCommentId(reportInfo.ReaderID, reportInfo.BookID, reportInfo.Date)
		if err != nil {
			return &common.LError{
				HttpCode: http.StatusInternalServerError,
				Msg:      "请求错误",
				Err:      errors.New("commentId错误"),
			}
		}
		// 更新Status
		err = commentRepository.UpdateStatusByCommentId(tx, commentId, 0)
		if err != nil {
			tx.Rollback()
			return &common.LError{
				HttpCode: http.StatusInternalServerError,
				Msg:      "请求错误",
				Err:      errors.New("更新Status错误"),
			}
		}
	}
	tx.Commit()
	return nil
}

func NewReportService() ReportService {
	return ReportService{
		DB: common.GetDB(),
	}
}
