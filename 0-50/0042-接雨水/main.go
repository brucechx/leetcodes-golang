package _042_接雨水

func trap(height []int) int {
	left, right := 0, len(height) - 1
	maxL, maxR := 0, 0
	var res int
	for left < right{
		maxL = max(maxL, height[left])
		maxR = max(maxR, height[right])
		if maxL < maxR{
			res += maxL - height[left]
			left++
		}else {
			res += maxR - height[right]
			right--
		}
	}
	return res
}

func max(a, b int) int{
	if a < b{
		return b
	}
	return a
}

