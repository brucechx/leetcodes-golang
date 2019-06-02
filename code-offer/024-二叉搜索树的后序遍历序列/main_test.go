package _24_二叉搜索树的后序遍历序列

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestVerifySequenceBst(t *testing.T) {
	cases := []struct{
		input []int
		output bool
	}{
		{[]int{5, 7, 6, 9, 10, 11, 8}, true},
		{[]int{7, 4, 6, 5}, false},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.output, VerifySequenceBst(cas.input))
	}
}
