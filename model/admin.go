package model

type Admin struct {
	Id       uint   `json:"id" gorm:"primaryKey;"`
	Phone    string `json:"phone" gorm:"not null;unique;"`
	Password string `json:"password" gorm:"default:null"`
}
