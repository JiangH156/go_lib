package model

type Book struct {
	BookId   string `json:"bookId" gorm:"type:varchar(50);primary_key;"`
	BookName string `json:"bookName" gorm:"type:varchar(20);"`
	Author   string `json:"author" gorm:"varchar(10);not null"`
	// 当前数量
	Amount uint `json:"amount"`
	// 位置
	Position string `json:"position" gorm:"type:varchar(30);"`
	// 总数量
	TotalAmount uint `json:"totalAmount"`
	// 借阅次数
	BorrowedTimes uint `json:"borrowedTimes"`
	// 状态
	Status int `json:"status"`
}
