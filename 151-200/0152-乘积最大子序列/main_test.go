package _152_乘积最大子序列

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProduct(t *testing.T){
	cases := []struct{
		input []int
		res int
	}{
		{
			input:[]int{2,3,-2,4},
			res: 6,
		},
		{
			input:[]int{-2,0,-1},
			res: 0,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, maxProduct(cas.input))
	}
}
