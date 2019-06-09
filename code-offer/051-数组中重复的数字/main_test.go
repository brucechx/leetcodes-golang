package _51_数组中重复的数字

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDuplicate(t *testing.T) {
	cases := []struct{
		input []int
		res int
	}{
		{[]int{2, 3, 1, 0, 2, 5, 3}, 2},
		{[]int{2, 1, 3, 1, 4}, 1},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, Duplicate1(cas.input))
		assert.Equal(t, cas.res, Duplicate2(cas.input))
		assert.Equal(t, cas.res, Duplicate3(cas.input))
	}

}
