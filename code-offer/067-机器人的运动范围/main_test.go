package _67_机器人的运动范围

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMovingCount(t *testing.T) {
	cases := []struct{
		threshold, rows, cols, res int
	}{
		{5, 10, 10, 21},
		{15, 20, 20, 359},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, MovingCount(cas.threshold, cas.rows, cas.cols))
	}
}
