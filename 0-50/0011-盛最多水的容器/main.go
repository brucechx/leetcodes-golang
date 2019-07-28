package _011_盛最多水的容器

func maxArea(height []int) int{
	// 从两端开始寻找，至少保证了宽度是最大值
	i, j := 0, len(height) - 1
	max := 0

	for i < j{
		h := min(height[i], height[j])
		area := h * (j - i)
		if max < area{
			max = area
		}

		if height[i] < height[j]{
			i++
		}else {
			j--
		}
	}
	return max
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}
