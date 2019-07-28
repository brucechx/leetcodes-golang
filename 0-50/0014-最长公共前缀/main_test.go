package _014_最长公共前缀

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T){
	cases := []struct{
		input []string
		res string
	}{
		{[]string{"flower","flow","flight"}, "fl"},
		{[]string{"dog","racecar","car"}, ""},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, longestCommonPrefix(cas.input))
	}
}
