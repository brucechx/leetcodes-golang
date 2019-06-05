package _45_孩子们的游戏_圆圈中最后剩下的数

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLastRemaining_Solution(t *testing.T) {
	cases := []struct{
		n, m int
		res int
	}{
		{1, 10, 0},
		{8, 5, 2},
		{6, 6, 3},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, LastRemainingSolution(cas.n, cas.m))
		assert.Equal(t, cas.res, LastRemainingSolution2(cas.n, cas.m))
	}
}
