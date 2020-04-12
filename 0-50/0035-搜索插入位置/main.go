package _035_搜索插入位置

func solution(nums []int, target int) int{
	left, right := 0, len(nums) - 1
	for left <= right{
		mid := left + (right - left) >> 1
		if nums[mid] == target{
			return mid
		}else if target > nums[mid]{
			left = mid + 1
		}else if target < nums[mid]{
			right = mid - 1
		}
	}
	return left
}
