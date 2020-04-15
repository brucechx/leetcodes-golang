package _033_搜索旋转排序数组

// 二分查找
func solution(nums []int, target int) int{
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := l + (r - l) / 2
		if nums[mid] == target{
			return mid
		}
		if nums[l] <= nums[mid]{ // 左边升序
			if target >= nums[l] && target <= nums[mid]{ // 左半边
				r = mid - 1
			}else {
				l = mid + 1
			}
		} else { // 右边升序
			if target >= nums[mid] && target <= nums[r]{
				l = mid + 1
			}else {
				r = mid - 1
			}
		}
	}
	return -1
}
