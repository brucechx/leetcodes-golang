package _034_在排序数组中查找元素的第一个和最后一个位置

// 二分查找，分别查找左右边界

func searchRange(nums []int, target int) []int {
	leftIndex := searchLeftIndex(nums, target)
	if leftIndex == -1{
		return []int{-1, -1}
	}
	rightIndex := searchRightIndex(nums, target)
	return []int{leftIndex, rightIndex}
}

func searchLeftIndex(nums []int, target int) int{
	left := 0
	right := len(nums) - 1
	for left <= right{
		mid := left + (right - left) >> 1
		if target == nums[mid]{
			right = mid - 1
		}else if target < nums[mid]{
			right = mid - 1
		}else if target > nums[mid]{
			left = mid + 1
		}
	}
	if left >= len(nums)  || nums[left] != target{
		return -1
	}
	return left
}

func searchRightIndex(nums []int, target int) int{
	left := 0
	right := len(nums) - 1
	for left <= right{
		mid := left + (right - left) >> 1
		if target == nums[mid]{
			left = mid + 1
		}else if target < nums[mid]{
			right = mid - 1
		}else if target > nums[mid]{
			left = mid + 1
		}
	}
	if right < 0  || nums[right] != target{
		return -1
	}
	return right
}