package model

type Admin struct {
	Id       uint   `json:"id" gorm:"primaryKey;unique;"`
	Phone    string `json:"phone" gorm:"not null;"`
	Password string `json:"password" gorm:"default:null"`
}
