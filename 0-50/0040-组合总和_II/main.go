package _040_组合总和_II

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	if len(candidates) == 0{
		return res
	}
	path := make([]int, 0)
	dfs(candidates, target, 0, path, &res) //深度优先
	return res
}

func dfs(candidates []int, target int, index int, path []int, res *[][]int){
	if target == 0{
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i:=index; i<len(candidates); i++{
		// 同层节点，数值相同
		if i != index && candidates[i] == candidates[i-1]{
			continue
		}
		// 剪枝
		if target - candidates[i] < 0{
			break
		}
		path = append(path, candidates[i])
		dfs(candidates, target-candidates[i], i+1, path, res)
		path = path[:len(path) - 1]
	}
}
