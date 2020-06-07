package _209_长度最小的子数组

// O(nlogn) 二分查找法
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sums := make([]int, n)
	sums[0] = nums[0]
	for i:=1; i<n; i++{
		sums[i] = nums[i] + sums[i-1]
	}

	res := 1 << 31 - 1
	for i:=0; i< n; i++{
		target := s - nums[i]
		index := binarySearch(i, n-1, sums, target + sums[i])
		if index != -1{
			res = min(res, index - i + 1)
		}
	}
	if res == 1 << 31 - 1{
		return 0
	}
	return res
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}

//寻求刚好大于 target 的 sums 的下标，也就是大于等于 target 所有 sums 中最小的那个
func binarySearch(start, end int, sums []int, target int) int{
	mid := -1
	for start <= end{
		mid = start + (end - start) >> 1
		if sums[mid] == target{
			return mid
		}else if sums[mid] < target{
			start = mid + 1
		}else {
			end = mid - 1
		}
	}
	if sums[mid] > target{
		return mid
	}
	return -1
}

