package _041_缺失的第一个正数

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T){
	cases := []struct{
		input []int
		res int
	}{
		{
			input:[]int{1, 2, 0},
			res: 3,
		},
		//{
		//	input:[]int{3, 4, -1, 1},
		//	res: 2,
		//},
		//{
		//	input:[]int{7,8,9,11,12},
		//	res:1,
		//},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, firstMissingPositive(cas.input))
	}
}
