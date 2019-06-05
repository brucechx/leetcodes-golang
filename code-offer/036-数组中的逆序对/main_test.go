package _36_数组中的逆序对

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInversePairs(t *testing.T) {
	cases := []struct{
		input []int
		res int
	}{
		{[]int{7, 5, 6, 4}, 5},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, InversePairs(cas.input))
	}
}
