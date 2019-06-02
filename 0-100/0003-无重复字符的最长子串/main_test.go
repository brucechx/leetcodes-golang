package _003_longest_substring_without_repeating_characters

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		input string
		output int
	}{
		{"ab", 2},
		{"abba", 2},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.output, LengthOfLongestSubstring(cas.input))
	}
}
