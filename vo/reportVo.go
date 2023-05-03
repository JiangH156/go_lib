package vo

import "Go_lib/model"

type ReportVo struct {
	Status       string     `json:"status"`       //状态
	CommentID    string     `json:"commentId"`    //评论id
	ReporterID   string     `json:"reporterId"`   //举报人id
	ReportDate   model.Time `json:"reportdate"`   //举报日期
	ReaderID     string     `json:"readerId"`     //读者id
	ReaderName   string     `json:"readerName"`   //读者姓名
	BookID       string     `json:"bookId"`       //书籍id
	BookName     string     `json:"bookName"`     //书名
	Date         model.Time `json:"date"`         //评论日期
	Content      string     `json:"content"`      //评论内容
	ReporterName string     `json:"reporterName"` //举报人姓名
}
