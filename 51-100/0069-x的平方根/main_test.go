package _069_x的平方根

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T)  {
	cases := []struct{
		input int
		res int
	}{
		{
			input: 10,
			res: 3,
		},
		{
			input: 4,
			res: 2,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, mySqrt(cas.input))
	}
}
