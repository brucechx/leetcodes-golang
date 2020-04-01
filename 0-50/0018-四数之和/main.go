package _018_四数之和

import (
	"sort"
)

/*
	同三数之和
	1. 先排序
    2. 使用四个指针(a<b<c<d); c:=b+1; d:=size-1
    3. 先固定，ab指针，然后cd包夹求解； 和偏大时，d--; 否则 c++
    4. b++ 进入b 循环
    5. a++ 进入a 循环

 */
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for a := 0; a < len(nums)-3; a++ {
		// 避免添加重复的结果
		// i>0 是为了防止nums[a-1]溢出
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		for b := a + 1; b < len(nums)-2; b++ {
			// 避免添加重复的结果
			// nums[b-1]虽然不会溢出
			// 但是不添加 b > a+1 的话，会漏掉那些a和b为相同数值的答案
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			c, d := b+1, len(nums)-1
			for c < d {
				s := nums[a] + nums[b] + nums[c] + nums[d]
				switch {
				case s < target:
					c++
				case s > target:
					d--
				default:
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					c, d = next(nums, c, d)
				}
			}
		}

	}
	return res
}

func next(nums []int, l, r int) (int, int) {
	// 所有满足条件且不重复的四元组
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}
