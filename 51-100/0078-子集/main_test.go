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
			input:[]int{1, 2},
		},
	}
	for _, cas := range cases{
		fmt.Println(subsets(cas.input))
	}
}
