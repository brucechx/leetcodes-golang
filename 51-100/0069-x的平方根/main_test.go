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


func TestMySqrtFloat32(t *testing.T)  {
	cases := []struct{
		input float32
		e float32
		res float32
	}{
		{
			input: float32(10),
			e: 0.01,
			res: 3.1542969,
		},
		{
			input: float32(4),
			e: 0.01,
			res: 1.9921875,
		},
	}
	for _, cas := range cases{
		//fmt.Println(mySqrtFloat(cas.input, cas.e))
		assert.Equal(t, cas.res, mySqrtFloat(cas.input, cas.e))
	}
}
