package _034_在排序数组中查找元素的第一个和最后一个位置

func solution(nums []int, target int) []int {
	if v := search(nums, target); v == -1 {
		return []int{-1, -1}
	}else {
		return getIndex(nums, v)
	}
}

func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right{
		mid := left + (right - left) / 2
		if nums[mid] == target{
			return mid
		}
		if target < nums[mid]{
			right = mid - 1
		}else {
			left = mid + 1
		}
	}
	return -1
}

func getIndex(nums []int, index int) []int{
	val := nums[index]
	l, r := index, index
	for {
		if l < 0 {
			break
		}
		if nums[l] != val{
			break
		}
		l--
	}
	for {
		if r >= len(nums){
			break
		}
		if nums[r] != val{
			break
		}
		r ++
	}
	return []int{l+1, r-1}
}