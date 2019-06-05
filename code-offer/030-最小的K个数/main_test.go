package _30_最小的K个数

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetLeastNumbersSolution(t *testing.T) {
	cases := []struct{
		input []int
		target int
		res []int
	}{
		{[]int{4,5,1,6,2,7,3,8}, 1, []int{1}},
		{[]int{4,5,1,6,2,7,3,8}, 4, []int{1, 2, 3, 4}},
		{[]int{4,5,1,6,2,7,3,8}, 10, []int{4,5,1,6,2,7,3,8}},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, GetLeastNumbersSolution(cas.input, cas.target))
	}
}
