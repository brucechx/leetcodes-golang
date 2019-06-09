package _52_构建乘积数组

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	cases := []struct{
		input []int
		res []int
	}{
		{[]int{1, 2, 3}, []int{6, 3, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, Multiply(cas.input))
	}
}
