package _31_连续子数组的最大和

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindGreatestSumOfSubArray(t *testing.T) {
	cases := []struct{
		input []int
		res int
	}{
		{[]int{6,-3,-2,7,-15,1,2,2}, 8},
		{[]int{-2}, -2},
		{[]int{-2, -1, -3}, -1},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, FindGreatestSumOfSubArray(cas.input))
		assert.Equal(t, cas.res, FindGreatestSumOfSubArray2(cas.input))
	}
}