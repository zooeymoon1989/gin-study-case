package main

import (
	"fmt"
	"gin-study-case/handlers/proto_buffers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


type TaskServer struct {
}

func main() {

	connections := os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_HOST")+"/"+os.Getenv("DB_DATABASE")+"?charset=utf8&parseTime=True&loc=Local"

	//var tasks TaskServer
	db, err := gorm.Open("mysql", connections)
	defer db.Close()
	if err != nil{
		fmt.Errorf("can not connect to mysql , err : %s" , err.Error())
	}
	//svr := grpc.NewServer()
	////todo.RegisterTasksServer(svr , tasks)
	//l , err := net.Listen("tcp","14444")
	//if err != nil {
	//	log.Fatalf("Can't listen tcp server with port is 14444 , err is :%v" , err.Error())
	//}
	//svr.Serve(l)
	r := gin.Default()
	r.POST("/todo/add",proto_buffers.Add)
	r.POST("/todo/list",proto_buffers.List)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200 , gin.H{
			"message":"pong",
		})
	})


	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
