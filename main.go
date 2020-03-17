package main

import (
	"fmt"
	"gin-study-case/routes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
)

func init() {
	//load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//默认不输出彩色日志
	gin.DisableConsoleColor()
	f,_ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)


}


func main() {

	connections := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + "/" + os.Getenv("DB_DATABASE") + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", connections)
	defer db.Close()
	if err != nil {
		fmt.Errorf("can not connect to mysql , err : %s", err.Error())
	}

	r :=routes.Route()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
