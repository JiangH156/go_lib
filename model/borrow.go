package model

// Borrow
// @Description:
type Borrow struct {
	Id       string `json:"id" gorm:"type:varchar(36);primaryKey"`
	ReaderId string `json:"readerId" gorm:"type:varchar(36);primaryKey;"`
	// Reader外键
	//Reader `gorm:"foreignKey:ReaderId;references:ReaderId"`
	BookId string `json:"bookId" gorm:"type:varchar(50);primaryKey"`
	// Book外键
	//Book `gorm:"foreignKey:BookId;references:BookId"`
	// 借阅日期
	BorrowDate Time `json:"borrowDate"`
	// 应归还日期
	ReturnDate Time `json:"ReturnDate"`
	// 实际归还日期
	RealData Time `json:"RealData"`
	// 借阅状态:借出中、已归还、逾期未还
	Status string `json:"status" gorm:"type:varchar(255)"`
}
