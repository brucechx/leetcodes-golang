package _213_打家劫舍II

// 转换为198题，两个单列最大值
func rob(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	if len(nums) == 0{
		return 0
	}
	return max(robb(nums[1:]), robb(nums[:len(nums)-1]))
}

func robb(nums []int) int{
	preMax, currMax := 0, 0
	for _, num := range nums{
		tmp := currMax
		currMax = max(preMax + num, currMax)
		preMax = tmp
	}
	return currMax
}

func max(a, b int) int{
	if a < b{
		return b
	}
	return a
}
