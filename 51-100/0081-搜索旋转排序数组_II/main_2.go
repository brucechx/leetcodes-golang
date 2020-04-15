package _081_搜索旋转排序数组_II

func search2(nums []int, target int) bool {
	if len(nums) == 0{
		return false
	}
	l, r := 0, len(nums) - 1
	for l <= r{
		mid := l + (r - l) >> 1
		if nums[mid] == target{
			return true
		}
		//
		if nums[mid] == nums[l]{
			l++
			continue
		}
		if nums[mid] > nums[l]{ // 前半部分有序
			if nums[mid] >= target && nums[l] <= target{
				r = mid - 1
			}else {
				l = mid + 1
			}
		}else { // 后半部分有序
			if nums[mid] <= target && nums[r] >= target{
				l = mid + 1
			}else {
				r = mid - 1
			}
		}
	}
	return false
}
