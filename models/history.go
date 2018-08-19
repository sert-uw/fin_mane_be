package models

import "github.com/jinzhu/gorm"

type History struct {
	gorm.Model
	Name       string `gorm:"not null"`
	Amount     int    `gorm:"not null"`
	CategoryID uint   `gorm:"not null"`
	AssetID    uint   `gorm:"not null"`
}
