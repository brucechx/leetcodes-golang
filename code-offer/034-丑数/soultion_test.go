package _34_丑数

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetUglyNumber(t *testing.T) {
	cases := []struct{
		input int
		res int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{10, 12},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, GetUglyNumberSolution1(cas.input))
		assert.Equal(t, cas.res, GetUglyNumberSolution2(cas.input))
	}
}
