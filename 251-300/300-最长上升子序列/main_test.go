package _00_最长上升子序列

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T){
	nums := []int{10, 9, 2, 5, 3, 7, 101, 8}
	fmt.Println(lengthOfLIS(nums))
}
