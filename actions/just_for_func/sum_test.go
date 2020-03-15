package just_for_func_test

import (
	"gin-study-case/actions/just_for_func"
	"testing"
)

func TestInts(t *testing.T)  {

	tt := []struct {
		name string
		numbers []int
		sum int
	}{
		{"one",[]int{1,2,3,4,5},15},
		{"two",nil,0},
		{"three",[]int{1,-1},0},
	}

	for _ , tc := range tt {
		t.Run(tc.name , func(t *testing.T) {
			s := just_for_func.Ints(tc.numbers)
			if s != tc.sum {
				t.Errorf("sum of %v should be %v ; got %v" , tc.name , tc.sum , s)
			}
		})
	}
}
