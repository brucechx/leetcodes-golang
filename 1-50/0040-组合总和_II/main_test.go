package _040_组合总和_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T){
	cases := []struct{
		input []int
		target int
		res [][]int
	}{
		{
			input:[]int{2, 3, 6, 7},
			target:7,
			res:[][]int{
				//{2, 2, 3},
				{7},
			},
		},
		{
			input:[]int{8, 7, 4, 3},
			target:11,
			res:[][]int{
				//{3, 4, 4},
				{3, 8},
				{4, 7},
			},
		},
	}
	for i, cas := range cases{
		if i != 1{
			continue
		}
		assert.Equal(t, cas.res, combinationSum(cas.input, cas.target))
	}
}