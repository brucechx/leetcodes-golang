package _004_寻找两个有序数组的中位数

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T){
	cases := []struct{
		nums1, nums2 []int
		res float64
	}{
		{[]int{1, 3}, []int{2}, float64(2)},
		{[]int{1, 3}, []int{2, 4}, float64(2.5)},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, findMedianSortedArrays(cas.nums1, cas.nums2))
	}
}
