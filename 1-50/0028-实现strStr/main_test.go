package _028_实现strStr

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T){
	cases := []struct{
		haystack, needle string
		res int
	}{
		{"hello", "ll", 2},
		{"aaaaaa", "bba", -1},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, strStr(cas.haystack, cas.needle))
	}

}
