package _046_全排列

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T){
	cases := []struct{
		input []int
		res [][]int
	}{
		//{
		//	input:[]int{1,2,3},
		//	res: [][]int{
		//		{1, 2, 3},
		//		{1, 3, 2},
		//		{2, 1, 3},
		//		{2, 3, 1},
		//		{3, 1, 2},
		//		{3, 2, 1},
		//	},
		//},
		{
			input:[]int{5,4,6,2},
		},
	}
	for _, cas := range cases{
		fmt.Println(permute3(cas.input))
	}
}
