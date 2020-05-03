package _162_寻找峰值

// 二分查找

func findPeakElement(nums []int) int {
	if len(nums) <= 1{
		return 0
	}
	l, r := 0, len(nums) - 1
	for l <= r{
		mid := l + (r - l) >> 1
		if mid - 1 >=0 && nums[mid] < nums[mid - 1]{
			r = mid - 1
		}else if mid + 1 < len(nums) && nums[mid] < nums[mid + 1]{
			l = mid + 1
		}else {
			return mid
		}
	}
	return l
}