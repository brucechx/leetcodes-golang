package _027_移除元素

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T){
	cases := []struct{
		input []int
		target int
		res int
		out  []int
	}{
		{[]int{1, 1, 2, 3}, 1, 2, []int{2, 3}},
		{[]int{1, 1, 2, 3}, 3, 3, []int{1, 1, 2}},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, removeElement(cas.input, cas.target))
		assert.Equal(t, cas.out, cas.input[:cas.res])
	}
}

func TestRemoveElement2(t *testing.T){
	cases := []struct{
		input []int
		target int
		res int
		out  []int
	}{
		{[]int{1, 1, 2, 3}, 1, 2, []int{3, 2}},
		{[]int{1, 1, 2, 3}, 3, 3, []int{1, 1, 2}},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, removeElement2(cas.input, cas.target))
		assert.Equal(t, cas.out, cas.input[:cas.res])
	}
}
