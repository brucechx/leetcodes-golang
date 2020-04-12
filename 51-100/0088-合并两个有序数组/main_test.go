package _088_合并两个有序数组

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T){
	m, n := 3, 3
	nums1 := make([]int, m+n)
	nums1[0], nums1[1], nums1[2] = 1, 2, 3
	nums2 := []int{2,5,6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1) // [1 2 2 3 5 6]
}
