package _38_数字在排序数组中出现的次数

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetNumberOfK2(t *testing.T) {
	cases := []struct{
		input []int
		target int
		res int
	}{
		{[]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 5, 3},
		{[]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 2, 1},
		{[]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 9, 2},
		{[]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 10, 0},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, GetNumberOfK2(cas.input, cas.target))
	}
}
