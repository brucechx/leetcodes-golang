package _32_从1到n整数中1出现的次数

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNumberOf1Between1AndN_Solution2(t *testing.T) {
	cases := []struct{
		input int
		res int
	}{
		{1, 1},
		{5, 1},
		{10, 2},
		{55, 16},
		{99, 20},
		{10000, 4001},
		{21345, 18821},
		{0, 0},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, NumberOf1Between1AndNSolution2(cas.input))
	}
}
