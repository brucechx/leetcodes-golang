package _011_盛最多水的容器

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T)  {
	cases := []struct{
		height []int
		res int
	}{
		{[]int{1,8,6,2,5,4,8,3,7}, 49},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, maxArea(cas.height))
	}
}
