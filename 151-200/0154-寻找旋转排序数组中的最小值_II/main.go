package _154_寻找旋转排序数组中的最小值_II

func findMin(nums []int) int {
	n := len(nums)
	i := 1
	for i < n && nums[i-1]<=nums[i]{
		i++
	}
	return nums[i%n]
}
