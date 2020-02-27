package _035_搜索插入位置

func solution(nums []int, target int) int{
	left, right := 0, len(nums) - 1
	mid := left + (right - left) / 2
	for left <= right{
		mid = left + (right - left) / 2
		if nums[mid] == target{
			return mid
		}
		if target > nums[mid]{
			left = mid + 1
		}else {
			right = mid - 1
		}
	}
	return left
}
