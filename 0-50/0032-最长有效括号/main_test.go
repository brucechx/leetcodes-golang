package _032_最长有效括号

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLongestValidParentheses(t *testing.T){
	cases := []struct{
		input string
		res int
	}{
		{"(()", 2},
		{")()())", 4},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, longestValidParentheses(cas.input))
	}
}
