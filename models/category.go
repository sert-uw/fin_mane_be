package models

type Category struct {
	Base
	Name       string `json:"name" gorm:"not null"`
	IsAddition bool   `json:"is_addition" gorm:"not null"`
}
