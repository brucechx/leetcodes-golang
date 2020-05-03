package _047_全排列_II

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T){
	cases := []struct{
		input []int
		res [][]int
	}{
		{
			input:[]int{1, 1, 2},
		},
	}
	for _, cas := range cases{
		//fmt.Println(permuteUnique(cas.input))
		fmt.Println(permuteUnique2(cas.input))
	}
}
