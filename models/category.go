package models

type Category struct {
	Base
	Name string `json:"name" gorm:"not null"`
	Type int    `json:"type" gorm:"not null"`
}
