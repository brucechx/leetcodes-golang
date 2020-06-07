package _220_存在重复元素III


func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	if len(nums) == 0{
		return false
	}

	n := len(nums)
	for l:=0; l<n-1; l++{
		for r:=l+1; r<n; r++{
			if r - l > k{
				break
			}
			if abs(nums[r] - nums[l]) <= t{
				return true
			}
		}
	}
	return false
}
