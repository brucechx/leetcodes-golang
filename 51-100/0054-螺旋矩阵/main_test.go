package _054_螺旋矩阵

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder(t *testing.T){
	cases := []struct{
		input [][]int
		res []int
	}{
		{
			input:[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			res: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, spiralOrder(cas.input))
	}
}
