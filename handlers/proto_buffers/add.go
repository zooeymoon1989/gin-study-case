package proto_buffers

import (
	"encoding/gob"
	"gin-study-case/protocal_buffers/todo"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"os"
)

func Add(c *gin.Context) {

	const dbPath = "./output/mydb.pb"
	var err error
	var todoProto = &todo.Task{}

	name := c.DefaultPostForm("name", "unknown")
	todoProto.Text = name
	todoProto.Done = false
	b, err := proto.Marshal(todoProto)
	if err != nil {
		c.JSON(200, err.Error())
		return
	}

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		c.JSON(200, err.Error())
		return
	}


	err = gob.NewEncoder(f).Encode(int32(len(b)))

	if err != nil {

		c.JSON(200, err.Error())
		return
	}

	_, err = f.Write(b)
	if err != nil {
		c.JSON(200, err.Error())
		return
	}

	if err = f.Close(); err != nil {
		c.JSON(200, err.Error())
		return
	}

	c.JSON(200, todoProto)
}
