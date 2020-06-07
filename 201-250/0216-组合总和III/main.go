package _216_组合总和III

// 回朔法
func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	helper(k, n, 1, &res, path)
	return res
}

func helper(k, n int, start int, res *[][]int, path []int){
	if len(path) == k && n == 0{
		tmp := make([]int, k)
		copy(tmp, path)
		*res = append(*res, tmp)
	}
	for i:=start; i<=9; i++{
		path = append(path, i)
		helper(k, n-i, i+1, res, path)
		path = path[:len(path)-1]
	}
}
