package _033_搜索旋转排序数组

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T){
	cases := []struct{
		input []int
		target int
		res int
	}{
		{
			input: []int{4,5,6,7,0,1,2},
			target: 0,
			res: 4,
		},
		{
			input: []int{4,5,6,7,0,1,2},
			target: 3,
			res: -1,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, solution(cas.input, cas.target))
	}
}
