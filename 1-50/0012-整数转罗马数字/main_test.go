package _012_整数转罗马数字

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T){
	cases := []struct{
		input int
		res string
	}{
		{3, "III"},
		{4, "IV"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, intToRoman(cas.input))
	}
}
