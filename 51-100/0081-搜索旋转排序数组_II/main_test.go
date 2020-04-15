package _081_搜索旋转排序数组_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T){
	cases := []struct{
		input []int
		target int
		res bool
	}{
		{
			input: []int{4,5,6,7,0,1,2},
			target: 0,
			res: true,
		},
		{
			input: []int{4,5,6,7,0,1,2},
			target: 3,
			res: false,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, search(cas.input, cas.target))
	}
}

