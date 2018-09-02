package controllers

import (
	"github.com/labstack/echo"
	"github.com/sert-uw/fin_mane_be/firebase"
	"net/http"
	"github.com/sert-uw/fin_mane_be/models"
	"github.com/sert-uw/fin_mane_be/db"
	"fmt"
	"strconv"
)

// 履歴一覧を取得する
func GetHistories(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token, err := firebase.GetUserToken(authHeader)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized."})
	}

	assetId := c.Param("asset_id")

	var asset models.Asset
	err = db.DB.Where("id = ? and user_id = ?", assetId, db.UserIdSubQuery(token.UID)).
		Preload("Histories").
		Find(&asset).
		Error

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Asset not found."})
	}

	return c.JSON(http.StatusOK, asset.Histories)
}

// 履歴を登録する
func PostHistory(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token, err := firebase.GetUserToken(authHeader)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized."})
	}

	assetId, err := strconv.Atoi(c.Param("asset_id"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request."})
	}

	var asset models.Asset
	isNoRecord := db.DB.Where("id = ? and user_id = ?", assetId, db.UserIdSubQuery(token.UID)).
		Find(&asset).RecordNotFound()
	if isNoRecord {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Asset Not Found."})
	}

	var history models.History
	if err := c.Bind(&history); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error."})
	}

	tx := db.DB.Begin()

	history.AssetID = uint(assetId)
	if err := tx.Create(&history).Preload("Category").First(&history).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error."})
	}

	if history.Category.IsAddition {
		asset.Balance += history.Amount
	} else {
		asset.Balance -= history.Amount
	}

	if err := tx.Save(&asset).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error."})
	}

	tx.Commit()

	return c.JSON(http.StatusCreated, history)
}
