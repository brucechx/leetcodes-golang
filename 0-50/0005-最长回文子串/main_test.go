package _005_最大回文数

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		input string
		output string
	}{
		{"babad", "aba"},
		{"cbbd", "bb"},
		{"a", "a"},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.output, longestPalindrome(cas.input))
	}
}
