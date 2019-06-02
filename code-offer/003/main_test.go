package _03

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	array := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	cases := []struct{
		input [][]int
		target int
		output bool
	}{
		{array, 7, true},
		{array, 12, true},
		{array, 5, false},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.output, Search(cas.input, cas.target))
	}
}
