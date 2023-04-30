package vo

import "Go_lib/model"

type BorrowVo struct {
	ReaderId   string     `json:"readerId"`
	BookId     string     `json:"bookId"`
	Status     string     `json:"status"`
	Author     string     `json:"author"`
	BookName   string     `json:"bookName"`
	BorrowDate model.Time `json:"borrowDate"`
	ReturnDate model.Time `json:"returnDate"`
	RealDate   model.Time `json:"realDate"`
	ReaderName string     `json:"readerName"`
}
