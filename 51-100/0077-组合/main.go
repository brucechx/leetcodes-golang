package _077_组合

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	if n <= 0 || k <= 0 {
		return res
	}
	path := make([]int, 0)
	dfs(n, k, 1, &res, path)
	return res
}

func dfs(n, k int, begin int, res *[][]int, path []int){
	if len(path) == k{
		tmp := make([]int, k)
		copy(tmp, path)
		*res = append(*res, tmp)
	}

	for i:=begin; i<=n; i++{
		path = append(path, i)
		dfs(n, k, i+1, res, path)
		path = path[:len(path) - 1]
	}
}


