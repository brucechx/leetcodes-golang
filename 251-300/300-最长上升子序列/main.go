package _00_最长上升子序列

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 || len(nums) == 1{
		return len(nums)
	}
	res := 1
	order := make([]int, 0)
	order = append(order, nums[0])
	for i:=1; i<len(nums); i++{
		if nums[i] > order[len(order)-1]{
			res++
			order = append(order, nums[i])
		}else {
			// 从已经排好的队列中找到第一个大于nums[i]的数进行替换
			// 二分查找
			target := nums[i]
			left, right := 0, len(order)
			for left < right{
				mid := left + (right - left) >> 1
				if order[mid] < target{
					left = mid + 1
				}else {
					right = mid
				}
			}
			order[left] = nums[i]
		}
	}
	return res
}
