package _029_两数相除

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T){
	cases := []struct{
		m, n int
		res int
	}{
		{10, 3, 3},
		{7, -3, -2},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, divide(cas.m, cas.n))
	}
}
