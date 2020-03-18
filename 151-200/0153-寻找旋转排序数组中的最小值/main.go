package _153_寻找旋转排序数组中的最小值

func findMin(nums []int) int {
	n := len(nums)
	i := 1
	for i < n && nums[i-1]<nums[i]{
		i++
	}
	return nums[i%n]
}

func findMin2(nums []int) int {
	left, right := 0, len(nums) - 1
	minIndex := left
	for nums[left] >= nums[right]{
		if right - left == 1{
			minIndex = right
			break
		}
		minIndex = (left + right) / 2
		if nums[left] == nums[right] && nums[minIndex] == nums[right]{
			return minOrder(nums, left, right)
		}

		if nums[minIndex] >= nums[left]{
			left = minIndex
		}else if nums[minIndex] <= nums[right]{
			right = minIndex
		}
	}
	return  nums[minIndex]
}

func minOrder(nums []int, l, r int) int{
	res := nums[l]
	for i:=l; i<=r; i++{
		if nums[i] < res{
			res = nums[i]
		}
	}
	return res
}
