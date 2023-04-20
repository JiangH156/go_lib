package vo

import "Go_lib/model"

type ReserveVO struct {
	ReaderId string     `json:"readerId"`
	BookId   string     `json:"bookId"`
	Status   string     `json:"status"`
	Author   string     `json:"author"`
	BookName string     `json:"bookName"`
	Date     model.Time `json:"date"`
}
