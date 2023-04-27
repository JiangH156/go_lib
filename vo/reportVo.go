package vo

import "Go_lib/model"

type ReportVo struct {
	Status       string     `json:"status"`       //状态
	CommentID    string     `json:"commentID"`    //评论id
	ReporterID   string     `json:"reporterID"`   //举报人id
	ReportDate   model.Time `json:"reportDate"`   //举报日期
	ReaderID     string     `json:"readerID"`     //读者id
	ReaderName   string     `json:"readerName"`   //读者姓名
	BookID       string     `json:"bookID"`       //书籍id
	BookName     string     `json:"bookName"`     //书名
	Date         model.Time `json:"date"`         //评论日期
	Content      string     `json:"content"`      //评论内容
	ReporterName string     `json:"reporterName"` //举报人姓名
}
