package _035_搜索插入位置

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T){
	cases := []struct{
		nums []int
		target int
		res int
	}{
		{
			nums: []int{1,3,5,6},
			target: 5,
			res: 2,
		},
		{
			nums: []int{1,3,5,6},
			target: 2,
			res: 1,
		},
		{
			nums: []int{1,3,5,6},
			target: 7,
			res: 4,
		},
		{
			nums: []int{1,3,5,6},
			target: 0,
			res: 0,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, solution(cas.nums, cas.target))
	}
}