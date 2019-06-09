package _54_表示数值的字符串

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsNumeric(t *testing.T) {
	cases := []struct{
		input string
		res bool
	}{
		{"+100", true},
		{"5e2", true},
		{"-123", true},
		{"3.1416", true},
		{"-1E-16", true},
		{"12e", false},
		{"1a3.14", false},
		{"1.2.3", false},
		{"+-5", false},
		{"12e+4.3", false},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, IsNumeric(cas.input))
	}

}