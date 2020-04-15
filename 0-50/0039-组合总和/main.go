package _039_组合总和

import "sort"

// 回溯法
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if len(candidates) == 0{
		return res
	}
	path := make([]int, 0)
	sort.Ints(candidates) // 需要排好序
	dfs(candidates, target, 0, path, &res)
	return res
}

func dfs(candidates []int, target int, begin int, path []int, res *[][]int){
	if target == 0{
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i:=begin; i< len(candidates); i++{
		// 剪枝
		if target - candidates[i] < 0{
			break
		}
		path = append(path, candidates[i])
		dfs(candidates, target-candidates[i], i, path, res)
		// 回溯
		path = path[:len(path)-1]
	}
}
