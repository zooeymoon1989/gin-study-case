package proto_buffers

import (
	"context"
	"fmt"
	"gin-study-case/protocal_buffers/todo_list"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
)

func List(c *gin.Context) {

	conn, err := grpc.Dial(":14431", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not dial to localhost:14431 , err: %v", err.Error())
	}
	client := todo_list.NewTasksClient(conn)

	results, err := client.List(context.Background(), &todo_list.Void{})
	if err != nil {
		fmt.Errorf("method list has problem , err: %v", err.Error())
	}


	c.JSON(200, results)
	//const dbPath = "./output/mydb.pb"
	//var todoProto = &todo.Task{}
	//var err error
	//
	//b, err := ioutil.ReadFile(dbPath)
	//if err != nil {
	//	c.JSON(200, err.Error())
	//	return
	//}
	//
	//for {
	//
	//	var length int32
	//
	//	if len(b) == 0 {
	//		c.JSON(200,"done")
	//		return
	//	}else if len(b) < 4 {
	//		c.JSON(200 , "bytes的长度小于4")
	//		return
	//	}
	//
	//	if err = gob.NewDecoder(bytes.NewReader(b[:4])).Decode(&length);err != nil {
	//		c.JSON(200 , err.Error())
	//		return
	//	}
	//
	//	b = b[4:]
	//
	//	if err = proto.Unmarshal(b[:length], todoProto); err != nil{
	//		c.JSON(200, err.Error())
	//		return
	//	}
	//
	//	b = b[length:]
	//
	//	if todoProto.Done {
	//		fmt.Println("no")
	//	}else{
	//		fmt.Println("yes")
	//	}
	//
	//	c.JSON(200 , todoProto.Text)
	//}

}
