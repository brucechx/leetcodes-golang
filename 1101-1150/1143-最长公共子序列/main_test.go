package _143_最长公共子序列

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T)  {
	cases := []struct{
		text1 string
		text2 string
		res int
	}{
		{
			text1: "abcde",
			text2: "ace",
			res:   3,
		},
		{
			text1: "abc",
			text2: "def",
			res:   0,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, longestCommonSubsequence(cas.text1, cas.text2))
	}
}
