package _047_全排列_II

import "sort"

func permuteUnique2(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[int]bool)
	dfs2(nums, path, &res, used)
	return res
}

func dfs2(nums []int, path []int, res *[][]int, used map[int]bool){
	if len(path) == len(nums){
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i:=0; i<len(nums);i++{
		if used[i]{
			continue
		}

		if i>0 && nums[i]==nums[i-1] && !used[i-1]{
			continue
		}

		used[i] = true
		path = append(path, nums[i])
		dfs2(nums, path, res, used)

		used[i] = false
		path = path[:len(path)-1]
	}
}
