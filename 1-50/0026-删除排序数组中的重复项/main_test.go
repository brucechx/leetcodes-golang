package _026_删除排序数组中的重复项

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T){
	cases := []struct{
		input []int
		res int
		out  []int
	}{
		{[]int{1, 1, 2, 3}, 3, []int{1, 2, 3}},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, removeDuplicates(cas.input))
		assert.Equal(t, cas.out, cas.input[:cas.res])
	}
}
