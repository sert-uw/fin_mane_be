package controllers

import (
	"github.com/labstack/echo"
	"github.com/sert-uw/fin_mane_be/firebase"
	"net/http"
	"github.com/sert-uw/fin_mane_be/models"
	"github.com/sert-uw/fin_mane_be/configs"
)

// Asset一覧を取得
func GetAssets(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token, err := firebase.GetUserToken(authHeader)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized."})
	}

	var user models.User
	if configs.DB.Preload("Assets").Where("token = ?", token.UID).First(&user).RecordNotFound() {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found."})
	}

	return c.JSON(http.StatusOK, user.Assets)
}

// Asset登録
func PostAsset(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token, err := firebase.GetUserToken(authHeader)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized."})
	}

	var user models.User
	if configs.DB.Preload("Assets").Where("token = ?", token.UID).First(&user).RecordNotFound() {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found."})
	}

	request := &struct {
		Name    string `json:"name"`
		Balance int    `json:"balance"`
	}{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad request."})
	}

	asset := models.Asset{Name: request.Name, Balance: request.Balance}
	user.Assets = append(user.Assets, asset)
	if err := configs.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Cannot insert asset."})
	}

	return c.JSON(http.StatusOK, user.Assets)
}

// Assert更新
func PutAsset(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token, err := firebase.GetUserToken(authHeader)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized."})
	}

	var user models.User
	if configs.DB.Preload("Assets").Where("token = ?", token.UID).First(&user).RecordNotFound() {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found."})
	}

	id := c.Param("id")
	var asset models.Asset
	if configs.DB.Where("id = ? and user_id = ?", id, user.ID).First(&asset).RecordNotFound() {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Asset not found."})
	}

	request := &struct {
		Name    string `json:"name"`
		Balance int    `json:"balance"`
	}{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad request."})
	}

	asset.Name = request.Name
	asset.Balance = request.Balance
	if err := configs.DB.Save(&asset).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Cannot update asset."})
	}

	return c.JSON(http.StatusOK, asset)
}