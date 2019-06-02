package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T){
	cases := []struct{
		input []int
		res [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, threeSum(cas.input))
	}
}
