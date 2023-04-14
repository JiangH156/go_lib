package model

// Borrow
// @Description:
type Borrow struct {
	Id       string `json:"id" gorm:"varchar(36);primary_key"`
	ReaderId string `json:"readerId" gorm:"varchar(50);primary_key"`
	BookId   string `json:"bookId" gorm:"type:varchar(50);primary_key"`
	// 借阅日期
	BorrowDate Time `json:"borrowDate"`
	// 应归还日期
	ReturnDate Time `json:"ReturnDate"`
	// 实际归还日期
	RealData Time `json:"RealData"`
	// 借阅状态:借出中、已归还、逾期未还
	Status string `json:"status" gorm:"varchar(255)"`
}
