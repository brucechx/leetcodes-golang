package _062_不同路径

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T){
	cases := []struct{
		m, n int
		res int
	}{
		{
			m:3,
			n:2,
			res:3,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, uniquePaths(cas.m, cas.n))
	}
}
