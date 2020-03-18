package _121_买卖股票的最佳时机

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
			input:[]int{7,1,5,3,6,4},
			res:5,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, maxProfit(cas.input))
	}
}
