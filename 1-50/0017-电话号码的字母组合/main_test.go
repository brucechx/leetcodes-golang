package _017_电话号码的字母组合

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T){
	cases := []struct{
		input string
		res []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, letterCombinations(cas.input))
	}
}
