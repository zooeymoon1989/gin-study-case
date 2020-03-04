package main

import (
	"gin-study-case/protocal_buffers/todo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type TaskServer struct {

}

func main() {

	var tasks TaskServer

	svr := grpc.NewServer()
	todo.RegisterTasksServer(svr , tasks)
	l , err := net.Listen("tcp","14444")
	if err != nil {
		log.Fatalf("Can't listen tcp server with port is 14444 , err is :%v" , err.Error())
	}
	svr.Serve(l)
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200 , gin.H{
	//		"message":"pong",
	//	})
	//})
	//
	//r.POST("/todo/add",proto_buffers.Add)
	//
	//
	//
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//r.GET("/floyd", floyd.Floyd)
	//r.Run("0.0.0.0:12345") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
