package vo

import "Go_lib/model"

type CommentVo struct {
	Email      string
	Status     uint
	CommentId  string
	ReaderId   string
	BookId     string
	ReaderName string
	BookName   string
	Date       model.Time
	Content    string
	Praise     uint
}
