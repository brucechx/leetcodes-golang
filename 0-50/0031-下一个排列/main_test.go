package _031_下一个排列

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T){
	cases := []struct{
		input []int
		res []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1, 2, 7, 4, 3, 1}, []int{1, 3, 1, 2, 4, 7}},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, nextPermutation(cas.input))
	}

}
