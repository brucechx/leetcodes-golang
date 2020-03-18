package _152_乘积最大子序列

func maxProduct(nums []int) int {
	if len(nums) == 0{
		return -1
	}
	curr, neg, maxRes := 1, 1, nums[0]
	for _, num := range nums{
		switch {
		case num > 0:
			curr, neg = curr * num, neg *num
		case num < 0:
			curr, neg = neg * num, curr *num
		default:
			curr, neg = 0, 1
		}
		if curr > maxRes{
			maxRes = curr
		}
		if curr <=0 {
			curr = 1
		}
	}
	return maxRes
}
