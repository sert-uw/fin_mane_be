package models

import "github.com/jinzhu/gorm"

type Asset struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Balance   int    `gorm:"not null; default:0"`
	UserID    uint   `gorm:"not null"`
	Histories []History
}
