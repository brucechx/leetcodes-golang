package _018_四数之和

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFourSum(t *testing.T){
	cases := []struct{
		input []int
		target int
		res [][]int
	}{
		{[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{{-2, -1, 1, 2}, {-2,  0, 0, 2}, {-1,  0, 0, 1}}},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, fourSum(cas.input, cas.target))
	}
}
