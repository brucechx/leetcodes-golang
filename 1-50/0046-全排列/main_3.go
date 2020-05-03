package _046_全排列


func permute3(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[int]bool) // 是否遍历过
	dfs3(nums, used, 0, path, &res)
	return res
}

func dfs3(nums []int, used map[int]bool,depth int, path []int, res *[][]int){
	if depth == len(nums){
		tmp := make([]int, depth)
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i:=0; i<len(nums); i++{
		if v, ok := used[i]; ok && v { // 这一层遍历过
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		dfs3(nums, used, depth+1, path, res)
		// 回溯
		used[i] = false
		path = path[:len(path)-1]
	}
}
