package main

import (
	"gin-study-case/handlers/proto_buffers"
	"gin-study-case/actions/floyd"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200 , gin.H{
			"message":"pong",
		})
	})

	r.POST("/todo/add",proto_buffers.Add)
	r.POST("/todo/list",proto_buffers.List)



	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.GET("/floyd", floyd.Floyd)
	r.Run("0.0.0.0:12345") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
