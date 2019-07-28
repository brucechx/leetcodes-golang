package _016_最接近的三数之和

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestThreeSumClosest(t *testing.T){
	cases := []struct{
		input []int
		target int
		res int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, threeSumClosest(cas.input, cas.target))
	}
}
