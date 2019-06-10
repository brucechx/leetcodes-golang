package _65_滑动窗口的最大值

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMaxInWindows(t *testing.T) {
	cases := []struct{
		input []int
		size int
		res []int
	}{
		{[]int{2,3,4,2,6,2,5,1}, 3, []int{4,4,6,6,6,5}},
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 10, []int{7}},
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, MaxInWindows(cas.input, cas.size))
	}
}
