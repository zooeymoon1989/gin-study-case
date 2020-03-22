package routes

import (
	"gin-study-case/actions/concurrecy_go"
	"gin-study-case/actions/trace"
	"gin-study-case/handlers/proto_buffers"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()
	r.Group("some")

	r.POST("/todo/add", proto_buffers.Add)
	r.POST("/todo/list", proto_buffers.List)
	r.POST("/cg/or_channel", concurrecy_go.OrChannel)
	r.POST("/error/error_handle",concurrecy_go.ErrorHandle)
	r.POST("/cg/pipeline",concurrecy_go.Pipeline)
	r.POST("/cg/repeat",concurrecy_go.Repeat)
	//trace工具
	r.POST("/forFunc/trace",trace.Trace)
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return r
}
