package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Token  string `gorm:"not null;unique"`
	Name   string `gorm:"not null"`
	Assets []Asset
}
