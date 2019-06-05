package _29_数组中出现次数超过一半的数字

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMoreThanHalfNumSolution(t *testing.T) {
	cases := []struct{
		input []int
		res int
	}{
		{[]int{1,2, 3, 2, 2, 2, 5, 4, 2}, 2},
		{[]int{1, 2}, 0},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, MoreThanHalfNumSolution(cas.input))
	}
}
