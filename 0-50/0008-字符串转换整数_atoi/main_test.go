package _008_字符串转换整数_atoi

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T){
	cases := []struct{
		input string
		res int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, myAtoi(cas.input))
	}
}
