package model

type Report struct {
	CommentId  string `json:"commentId" gorm:"type:varchar(50);primary_key"`
	ReporterId string `json:"reporterId" gorm:"type:varchar(50);primary_key"`
	ReportDate Time   `json:"ReportDate" gorm:"type:varchar(50);primary_key"`
	Status     string `json:"status"`
}
