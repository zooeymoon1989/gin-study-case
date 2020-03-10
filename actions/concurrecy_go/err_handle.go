package concurrecy_go

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Err error
	Response *http.Response
}

func ErrorHandle(c *gin.Context)  {

	checkStatus := func(done <-chan interface{} , urls... string) <-chan Result{
		results := make(chan Result)
		go func() {
			defer close(results)
			for _ , url := range urls {
				resp , err := http.Get(url)
				result := Result{err,resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.baidu.com","https://badhost"}
	for response := range checkStatus(done,urls...) {
		if response.Err != nil {
			fmt.Printf("Error : %v\n",response.Err)
			continue
		}
		fmt.Printf("Response: %v\n",response.Response.Status)
	}

}


