package _038_外观数列

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAndSay(t *testing.T){
	cases := []struct{
		input int
		res string
	}{
		{
			input: 4,
			res: "1211",
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, countAndSay(cas.input))
	}
}