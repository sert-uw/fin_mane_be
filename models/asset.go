package models

type Asset struct {
	Base
	Name      string    `json:"name" gorm:"not null"`
	Balance   int       `json:"balance" gorm:"not null; default:0"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Histories []History `json:"histories"`
}
