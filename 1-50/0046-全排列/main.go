package _046_全排列

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	dfs(nums, track, 0, &res)
	return res
}

func dfs(nums []int, path []int, start int, res *[][]int){
	if start == len(nums){
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	for i:=start; i<len(nums); i++{
		// 以i为界划分，左右两部分
		nums[start], nums[i] = nums[i], nums[start]
		dfs(nums, path, start+1, res)
		// 回溯
		nums[start], nums[i] = nums[i], nums[start]
	}
}
