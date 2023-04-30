package repository

import (
	"Go_lib/common"
	"Go_lib/model"
	"Go_lib/vo"
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
		Model(&model.Report{}).
		Select(`reports.*,books.book_name,c.reader_id, c.book_id,c.date, c.content, readers.reader_name AS reportName`).
		Joins("LEFT JOIN comments c ON c.comment_id = reports.comment_id").
		Joins("LEFT JOIN books ON books.book_id = c.book_id").
		Joins("LEFT JOIN readers ON readers.reader_id = c.reader_id").
		Where("readers.reader_id = reports.reporter_id AND readers.reader_id = ?", readerId).
		Find(&reports).Error; err != nil {
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
		Model(&model.Report{}).
		Select(`reports.*,readers.email, books.book_name,c.reader_id, c.book_id,c.date, c.content, readers.reader_name AS reportName`).
		Joins("LEFT JOIN comments c ON c.comment_id = reports.comment_id").
		Joins("LEFT JOIN books ON books.book_id = c.book_id").
		Joins("LEFT JOIN readers ON readers.reader_id = c.reader_id").
		Where("readers.reader_id = reports.reporter_id").
		Scan(&reportVos).Error; err != nil {
		return reportVos, err
	}
	return reportVos, nil
}

func NewReportRepository() ReportRepository {
	return ReportRepository{
		DB: common.GetDB(),
	}
}
