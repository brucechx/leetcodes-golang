package _35_第一个只出现一次的字符位置

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFirstNotRepeatingChar(t *testing.T) {
	cases := []struct{
		input string
		res int
	}{
		{"hello", 0},
		{"hhLLo", 4},
		{"hhLlo", 2},
		{"hhLLoo", -1},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, FirstNotRepeatingChar(cas.input))
	}
}
