package _090_子集_II

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(nums)
	dfs(nums, path, 0, &res)
	return res
}

func dfs(nums []int, path []int, pos int, res *[][]int){
	tmp := make([]int, len(path))
	copy(tmp, path)
	*res = append(*res, tmp)

	for i:=pos; i<len(nums); i++{
		// 剪枝
		if i != pos && nums[i] == nums[i-1]{
			continue
		}
		path = append(path, nums[i])
		dfs(nums, path, i+1, res)
		path = path[:len(path) - 1]
	}
}
