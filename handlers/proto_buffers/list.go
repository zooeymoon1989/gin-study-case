package proto_buffers

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"gin-study-case/protocal_buffers/todo"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

func List(c *gin.Context) {
	const dbPath = "./output/mydb.pb"
	var todoProto = &todo.Task{}
	var err error

	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		c.JSON(200, err.Error())
		return
	}

	for {

		var length int32

		if len(b) == 0 {
			c.JSON(200,"done")
			return
		}else if len(b) < 4 {
			c.JSON(200 , "bytes的长度小于4")
			return
		}

		if err = gob.NewDecoder(bytes.NewReader(b[:4])).Decode(&length);err != nil {
			c.JSON(200 , err.Error())
			return
		}

		b = b[4:]

		if err = proto.Unmarshal(b[:length], todoProto); err != nil{
			c.JSON(200, err.Error())
			return
		}

		b = b[length:]

		if todoProto.Done {
			fmt.Println("no")
		}else{
			fmt.Println("yes")
		}

		c.JSON(200 , todoProto.Text)
	}





}
