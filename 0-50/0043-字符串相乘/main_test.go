package _043_字符串相乘

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T){
	cases := []struct{
		input1 string
		input2 string
		res string
	}{
		{
			input1:"2",
			input2:"3",
			res:"6",
		},
		{
			input1:"123",
			input2:"45",
			res:"5535",
		},
		{
			input1:"123",
			input2:"456",
			res:"56088",
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, multiply(cas.input1, cas.input2))
	}
}
