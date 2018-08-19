package main

import (
	"github.com/sert-uw/fin_mane_be/configs"
	"fmt"
)

func main() {
	// DB接続
	if err := configs.Init(); err != nil {
		fmt.Println(err)
		return
	}
}