package _41_和为S的连续正数序列

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindContinuousSequence(t *testing.T) {
	cases := []struct{
		input int
		res [][]int
	}{
		{5, [][]int{{2, 3}}},
		{100, [][]int{{9,10, 11, 12, 13, 14, 15, 16}, {18, 19, 20, 21, 22}}},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, FindContinuousSequence(cas.input))
	}
}
