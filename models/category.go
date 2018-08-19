package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"not null"`
	Type int    `gorm:"not null"`
}
