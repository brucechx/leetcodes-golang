package _44_扑克牌顺子

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsContinuous(t *testing.T) {
	cases := []struct{
		input []int
		res bool
	}{
		{[]int{3, 5, 1, 0, 4}, true},
		{[]int{3, 5, 4, 7, 6}, true},
		{[]int{3, 5, 7, 4, 8}, false},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, IsContinuous(cas.input))
	}
}
