package controllers

import (
	"github.com/labstack/echo"
	"github.com/sert-uw/fin_mane_be/firebase"
	"net/http"
	"github.com/sert-uw/fin_mane_be/models"
	"github.com/sert-uw/fin_mane_be/db"
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
