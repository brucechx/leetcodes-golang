package _49_把字符串转换成整数

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStrToInt(t *testing.T) {
	cases := []struct{
		input string
		res int
	}{
		{"1a33", 0},
		{"-2147483648", -2147483648},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, StrToInt(cas.input))
	}
}
