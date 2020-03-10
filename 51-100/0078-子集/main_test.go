package _078_子集

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
			input:[]int{1, 2, 3},
			res: [][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 3},
				{1, 3},
				{2},
				{2, 3},
				{3},
			},
		},
	}
	for _, cas := range cases{
		fmt.Println(subsets(cas.input))
		//fmt.Println(subset2(cas.input))
	}
}
