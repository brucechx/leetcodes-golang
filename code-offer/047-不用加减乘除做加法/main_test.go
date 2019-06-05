package _47_不用加减乘除做加法

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T){
	cases := []struct{
		a, b int
		res int
	}{
		{3, 4, 7},
		{1, 2, 3},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, Add1(cas.a, cas.b))
		assert.Equal(t, cas.res, Add2(cas.a, cas.b))
	}
}
