package main

import (
	"fmt"

	"github.com/daitasu/test-module/db"
	"github.com/daitasu/test-module/routes"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	fmt.Println("実行！！！！")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	db.Init()
	routes.Init()
	db.Close()
}
