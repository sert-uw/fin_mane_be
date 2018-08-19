package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"os"
	"github.com/sert-uw/fin_mane_be/models"
)

var (
	dbConf = fmt.Sprintf("%s:%s@tcp(%s:%s)/fin_mane?parseTime=true&loc=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_IP"), os.Getenv("DB_PORT"), "Asia%2FTokyo")
	DB *gorm.DB
)

// DB接続の初期化を行う
func Init() error {
	var err error

	if DB, err = gorm.Open("mysql", dbConf); err != nil {
		return err
	}

	if err := DB.AutoMigrate(&models.User{}, &models.Asset{}, &models.Category{}, &models.History{}).Error; err != nil {
		return err
	}

	return nil
}
