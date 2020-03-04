package floyd

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Floyd(c *gin.Context) {
	var intArray = []int{1, 3, 2, 5, 4, 5, 6, 7, 8, 9}
	var tortoise, hare, ptr1 = intArray[0], intArray[0], intArray[0]
	var ptr2 int
	for {
		tortoise = intArray[tortoise]
		hare = intArray[intArray[hare]]
		if tortoise == hare {
			break
		}
	}

	ptr2 = tortoise

	for ptr1 != ptr2 {
		ptr1 = intArray[ptr1]
		ptr2 = intArray[ptr2]
	}

	fmt.Println(ptr1)

}
