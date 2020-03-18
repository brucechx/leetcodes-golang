package _047_全排列_II

import (
	"sort"
)

func permuteUnique2(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	dfs(nums, 0, &res)
	return res
}

func dfs(nums []int, start int,  res *[][]int){
	if start == len(nums){
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
	}
	exist := make(map[int]bool)
	for i:=start; i<len(nums); i++{
		if exist[i]{
			continue
		}
		exist[i] = true
		nums[start], nums[i] = nums[i], nums[start]
		dfs(nums, start+1, res)
		nums[start], nums[i] = nums[i], nums[start]
	}
}
