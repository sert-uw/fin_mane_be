package models

type User struct {
	Base
	Token  string  `json:"token" gorm:"not null;unique"`
	Name   string  `json:"name" gorm:"not null"`
	Assets []Asset `json:"assets"`
}
