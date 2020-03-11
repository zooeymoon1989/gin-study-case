package concurrecy_go

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Pipeline(c *gin.Context) {

	//生成器
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:

				}
			}
		}()

		return intStream
	}

	multiply := func(done <-chan interface{} , intStream <-chan int , multiplier int) <-chan int{
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream{
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func(done chan interface{} , intStream <-chan int , additive int)<-chan int{
		addStream := make(chan int)
		go func() {
			defer close(addStream)
			for i := range intStream{
				select {
				case <-done:
					return
				case addStream <- i + additive:
				}
			}
		}()
		return addStream
	}
	start := time.Now()

	done := make(chan interface{})
	intStream := generator(done,1,2,3,4,5,6,7,8,9)
	pipeline := multiply(done,add(done , multiply(done ,intStream,2),1),3)
	for i := range pipeline {
		fmt.Println(i)
	}
	c.JSON(200 , time.Since(start).Seconds())
}
