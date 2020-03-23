package unsafe_pointer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"unsafe"
)

func UnsafePointer(c *gin.Context)  {
	i := 10
	ip := &i
	fmt.Println(i)
	fmt.Println(ip)
	var fp  = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(fp)
	fmt.Println(i)
}
