package main

import (
	"github.com/labstack/echo"
	"github.com/sert-uw/fin_mane_be/db"
	"fmt"
	"net/http"
	"os"
	"github.com/sert-uw/fin_mane_be/controllers"
	"github.com/sert-uw/fin_mane_be/firebase"
)

func main() {
	// DB接続
	if err := db.Init(); err != nil {
		fmt.Println(err)
		return
	}

	// Firebaseのセットアップ
	if err := firebase.Init(); err != nil {
		fmt.Println(err)
		return
	}

	// Echo設定
	e := echo.New()
	e.Use(firebase.JWTHandler)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
	})

	e.GET("/categories", controllers.GetCategories)
	e.POST("/categories", controllers.PostCategory)
	e.PUT("/categories/:id", controllers.PutCategory)
	e.DELETE("/categories/:id", controllers.DeleteCategory)

	e.GET("/user", controllers.GetUser)
	e.PUT("/user", controllers.PutUser)

	e.GET("/assets", controllers.GetAssets)
	e.POST("/assets", controllers.PostAsset)
	e.PUT("/assets/:id", controllers.PutAsset)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
