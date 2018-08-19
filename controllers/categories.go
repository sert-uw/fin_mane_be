package controllers

import (
	"github.com/labstack/echo"
	"github.com/sert-uw/fin_mane_be/configs"
	"github.com/sert-uw/fin_mane_be/models"
	"fmt"
	"net/http"
)

// カテゴリ一覧を返す
func GetCategories(c echo.Context) error {
	var categories []models.Category
	if err := configs.DB.Find(&categories).Error; err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, categories)
}

// カテゴリを登録する
func PostCategory(c echo.Context) error {
	var category models.Category
	if err := c.Bind(&category); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request."})
	}

	if err := configs.DB.Create(&category).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error."})
	}

	return c.JSON(http.StatusCreated, category)
}

// カテゴリを更新する
func PutCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	if configs.DB.Where("id = ?", id).Find(&category).RecordNotFound() {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Not Found."})
	}

	if err := c.Bind(&category); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request."})
	}

	if err := configs.DB.Save(&category).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error."})
	}

	return c.JSON(http.StatusOK, category)
}

// カテゴリを削除する
func DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	if configs.DB.Where("id = ?", id).Find(&category).RecordNotFound() {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Not Found."})
	}

	if err := configs.DB.Delete(&category).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error."})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}
