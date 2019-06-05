package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSumN(t *testing.T)  {
	cases := []struct{
		input int
		res int
	}{
		{3, 6},
		{5, 15},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, SumN(cas.input))
	}
}



