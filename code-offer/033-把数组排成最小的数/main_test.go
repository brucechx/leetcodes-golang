package _33_把数组排成最小的数

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPrintMinNums(t *testing.T) {
	cases := []struct{
		input []int
		res string
	}{
		{[]int{3, 32, 321}, "321323"},
		{[]int{1}, "1"},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, PrintMinNums(cas.input))
	}
}
