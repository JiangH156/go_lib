package repository

import (
	"github.com/jiangh156/Go_lib/common"
	"github.com/jiangh156/Go_lib/model"
	"github.com/jiangh156/Go_lib/vo"
	"gorm.io/gorm"
)

type ReportRepository struct {
	DB *gorm.DB
}

// GetReportRecordsByReaderId
// @Description  获取举报记录
// @Author John 2023-04-26 19:31:05
// @Param readerId
// @Return reports
// @Return lErr
func (r *ReportRepository) GetReportRecordsByReaderId(readerId string) (reports []vo.ReportVo, err error) {
	if err = r.DB.
		Debug().
		Table("reports").
		Select(`reports.*,books.book_name,c.reader_id, c.book_id,c.date, c.content, readers.reader_name,readers.reader_name AS ReporterName`).
		Joins("LEFT JOIN comments c ON c.comment_id = reports.comment_id").
		Joins("LEFT JOIN readers ON readers.reader_id = c.reader_id").
		Joins("LEFT JOIN books ON books.book_id = c.book_id").
		Where("readers.reader_id = ?", readerId).
		Scan(&reports).
		Error; err != nil {
		return reports, err
	}
	return reports, nil
}

// GetAllReportRecords
// @Description 返回所有举报记录
// @Author John 2023-04-28 15:13:11
// @Return reports
// @Return err
func (r *ReportRepository) GetAllReportRecords() (reportVos []vo.ReportVo, err error) {
	if err = r.DB.
		Table("reports").
		Select(`reports.*,readers.email, books.book_name,c.reader_id, c.book_id,c.date, c.content, readers.reader_name,readers.reader_name AS ReporterName`).
		Joins("LEFT JOIN comments c ON reports.comment_id = c.comment_id").
		Joins("LEFT JOIN books ON c.book_id = books.book_id").
		Joins("LEFT JOIN readers ON c.reader_id = readers.reader_id").
		//Where("readers.reader_id = reports.reporter_id").
		Scan(&reportVos).Error; err != nil {
		return reportVos, err
	}
	return reportVos, nil
}

// CreateReport
// @Description 用户举报评论
// @Author John 2023-05-02 13:43:21
// @Param tx
// @Param commentId
// @Param reporterId
// @Return error
func (r *ReportRepository) CreateReport(tx *gorm.DB, report model.Report) error {
	if err := tx.Create(report).Error; err != nil {
		return err
	}
	return nil
}

// UpdateStatus
// @Description 驳回举报
// @Author John 2023-05-03 15:30:32
// @Param tx
// @Param report
// @Return error
func (r *ReportRepository) UpdateStatus(tx *gorm.DB, report model.Report) error {
	if err := tx.
		Model(&model.Report{}).
		Where("comment_id = ? AND reporter_id = ? AND report_date = ?", report.CommentId, report.ReporterId, report.ReportDate).
		UpdateColumn("status", report.Status).
		Error; err != nil {
		return err
	}
	return nil
}

func NewReportRepository() ReportRepository {
	return ReportRepository{
		DB: common.GetDB(),
	}
}
