package _136_只出现一次的数字

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T){
	cases := []struct{
		input []int
		res int
	}{
		{
			input:[]int{2, 2, 1},
			res: 1,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, singleNumber(cas.input))
	}
}
