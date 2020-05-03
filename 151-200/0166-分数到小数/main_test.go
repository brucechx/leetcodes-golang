package _166_分数到小数

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T){
	cases := []struct{
		numerator int
		denominator int
		res string
	}{
		{
			numerator: 1,
			denominator: 2,
			res:"0.5",
		},
		{
			numerator: 1,
			denominator:3,
			res:"0.(3)",
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, fractionToDecimal(cas.numerator, cas.denominator))
	}

}
