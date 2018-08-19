package main

import (
	"github.com/labstack/echo"
	"github.com/sert-uw/fin_mane_be/configs"
	"fmt"
	"net/http"
	"os"
	"github.com/sert-uw/fin_mane_be/controllers"
)

func main() {
	// DB接続
	if err := configs.Init(); err != nil {
		fmt.Println(err)
		return
	}

	// Echo設定
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
	})

	e.GET("/categories", controllers.GetCategories)
	e.POST("/categories", controllers.PostCategory)
	e.PUT("/categories/:id", controllers.PutCategory)
	e.DELETE("/categories/:id", controllers.DeleteCategory)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
