package main

import (
	"gin-study-case/actions/floyd"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/floyd", floyd.Floyd)
	r.Run("0.0.0.0:12345") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
