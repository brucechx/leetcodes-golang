package _016_最接近的三数之和

import (
	"math"
	"sort"
)

/*
	排序 & 双指针法
*/

func threeSumClosest(nums []int, target int) int{
	// 排序
	sort.Ints(nums)
	res, delta := 0, math.MaxInt64

	for i := range nums{
		// 避免重复计算
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 双指针
		l, r := i+1, len(nums)-1
		for l < r{
			s := nums[i] + nums[l] + nums[r]
			switch {
			// 因为已经排序，和小则左子针进位
			case s < target:
				l++
				if delta > target - s{
					delta = target - s
					res = s
				}
			case s > target:
				r--
				if delta > s - target{
					delta = s - target
					res = s
				}
			default:
				return s
			}
		}
	}
	return res
}
