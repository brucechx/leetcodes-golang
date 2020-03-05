package _090_子集_II

import (
	"fmt"
	"testing"
)

func TestSubSet(t *testing.T){
	cases := []struct{
		input []int
		res [][]int
	}{
		{
			input:[]int{1, 2, 2},
		},
		{
			input:[]int{4, 4, 1, 4},
		},
	}
	for _, cas := range cases{
		fmt.Println(subsetsWithDup(cas.input))
	}
}

