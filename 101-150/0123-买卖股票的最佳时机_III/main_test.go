package _123_买卖股票的最佳时机_III

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T){
	cases := []struct{
		input []int
		res int
	}{
		{
			input:[]int{3,3,5,0,0,3,1,4},
			res:6,
		},
		//{
		//	input:[]int{7,6,4,3,1},
		//	res:0,
		//},
		//{
		//	input:[]int{1,2,3,4,5},
		//	res:4,
		//},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, maxProfit(cas.input))
	}
}
