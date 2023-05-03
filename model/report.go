package model

type Report struct {
	CommentId  string `json:"commentId" gorm:"type:varchar(50);primary_key"`
	ReporterId string `json:"reporterId" gorm:"type:varchar(50);primary_key"`
	ReportDate Time   `json:"ReportDate" gorm:"primary_key"`
	// 审核中、已驳回、已通过
	Status string `json:"status" gorm:"type:varchar(50)"`
}
