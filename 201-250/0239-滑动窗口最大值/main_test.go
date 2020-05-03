package _239_滑动窗口最大值

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T)  {
	cases := []struct{
		input []int
		k int
		res []int
	}{
		{
			input:[]int{9,10,9,-7,-4,-8,2,-6},
			k:5,
			res:[]int{10, 10, 9, 2},
		},
		{
			input:[]int{1, -1},
			k:1,
			res: []int{1, -1},
		},
		{
			input:[]int{1, 2, 3},
			k:3,
			res:[]int{3},
		},
		{
			input:[]int{1,3,-1,-3,5,3,6,7},
			k:3,
			res:[]int{3,3,5,5,6,7},
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, maxSlidingWindow(cas.input, cas.k))
		//fmt.Println(maxSlidingWindow(cas.input, cas.k))
	}
}
