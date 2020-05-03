package _047_全排列_II

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	dfs(nums, 0, &res)
	return res
}

func dfs(nums []int, depth int, res *[][]int){
	if depth == len(nums){
		tmp := make([]int, depth)
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	exited := make(map[int]bool)
	for i:=depth; i<len(nums); i++{
		if exited[nums[i]]{
			continue
		}
		exited[nums[i]] = true
		nums[depth], nums[i] = nums[i], nums[depth]
		dfs(nums, depth+1, res)
		nums[depth], nums[i] = nums[i], nums[depth]
	}
}
