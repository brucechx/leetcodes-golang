package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		input []int
		target int
		expect []int
	}{
		{[]int{3, 4, 5, 6}, 7, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 8, []int{2, 4}},
		{[]int{1, 2, 3, 4, 5}, 10, []int{}},
	}

	for _, cas := range cases{
		// 方法1
		res1 := Solution1(cas.input, cas.target)
		assert.Equal(t, cas.expect, res1)

		// 方法2
		res2 := Solution2(cas.input, cas.target)
		assert.Equal(t, cas.expect, res2)

		// 方法3
		res := Solution(cas.input, cas.target)
		assert.Equal(t, cas.expect, res )
	}
}
