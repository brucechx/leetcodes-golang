package _053_最大子序和

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T){
	cases := []struct{
		input []int
		res int
	}{
		{
			input:[]int{-2,1,-3,4,-1,2,1,-5,4},
			res: 6,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, maxSubArray(cas.input))
	}
}
