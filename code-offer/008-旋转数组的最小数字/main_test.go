package _08_旋转数组的最小数字

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	cases := []struct{
		input []int
		target int
	}{
		{[]int{1, 2, 3, 4}, 1},  // 场景一
		{[]int{3, 4, 5, 1, 2}, 1}, // 场景二： 常规
		{[]int{1, 0, 1, 1, 1}, 0}, // 场景三： 特殊
	}
	for _, cas := range cases{
		//got := FindMin(cas.input)
		got := findMin(cas.input)
		assert.Equal(t, cas.target, got)
	}
}
