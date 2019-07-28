package _022_括号生成

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T){
	cases := []struct{
		input int
		res []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, generateParenthesis(cas.input))
	}
}
