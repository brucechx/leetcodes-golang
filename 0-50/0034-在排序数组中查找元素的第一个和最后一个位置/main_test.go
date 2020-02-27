package _034_在排序数组中查找元素的第一个和最后一个位置

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T)  {
	cases := []struct{
		input []int
		target int
		res []int
	}{
		{
			input: []int{5,7,7,8,8,10},
			target: 8,
			res: []int{3, 4},
		},
		{
			input: []int{5,7,7,8,8,10},
			target: 6,
			res: []int{-1, -1},
		},
		{
			input: []int{2, 2},
			target: 2,
			res: []int{0, 1},
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, solution(cas.input, cas.target))
	}
}