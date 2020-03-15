package just_for_func_test

import (
	"fmt"
	"gin-study-case/actions/just_for_func"
)

func ExampleInts() {
	s := just_for_func.Ints([]int{1,2,3,4,5})
	fmt.Println("sum of one to five is",s)
	// Output:
	// sum of one to five is 15
}
