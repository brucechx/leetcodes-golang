package _042_接雨水

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T){
	cases := []struct{
		input []int
		res int
	}{
		{
			input:[]int{0,1,0,2,1,0,1,3,2,1,2,1},
			res: 6,
		},
	}
	for _, cas:= range cases{
		assert.Equal(t, cas.res, trap(cas.input))
	}
}