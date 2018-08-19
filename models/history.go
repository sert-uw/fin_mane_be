package models

type History struct {
	Base
	Name       string `json:"name" gorm:"not null"`
	Amount     int    `json:"amount" gorm:"not null"`
	CategoryID uint   `json:"category_id" gorm:"not null"`
	AssetID    uint   `json:"asset_id" gorm:"not null"`
}
