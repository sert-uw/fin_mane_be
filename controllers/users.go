package controllers

import (
	"github.com/labstack/echo"
	"github.com/sert-uw/fin_mane_be/firebase"
	"net/http"
	"github.com/sert-uw/fin_mane_be/configs"
	"github.com/sert-uw/fin_mane_be/models"
)

// ユーザー情報取得
func GetUser(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token, err := firebase.GetUserToken(authHeader)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	var user models.User
	if configs.DB.Preload("Assets").Where("token = ?", token.UID).First(&user).RecordNotFound() {
		name, err := firebase.GetUserName(token.UID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Cannot get user name."})
		}

		user = models.User{Token: token.UID, Name: name}
		if err := configs.DB.Preload("Assets").Create(&user).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Cannot insert user."})
		}
	}

	return c.JSON(http.StatusOK, user)
}