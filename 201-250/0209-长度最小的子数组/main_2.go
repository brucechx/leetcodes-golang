package _209_长度最小的子数组

// 双指针法 O(n)
func minSubArrayLen2(s int, nums []int) int {
	n := len(nums)
	if n == 0{
		return 0
	}
	left, right := 0, 0
	sum := 0
	res := 1 << 31 - 1
	flag := false
	for ;right < n; right++ {
		sum += nums[right]
		for sum >= s {
			flag = true
			if right - left + 1 < res{
				res = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if ! flag{
		return 0
	}
	return res
}
