package _189_旋转数组

func rotate(nums []int, k int)  {
	k %= len(nums)
	// 第一反转
	reverse(nums, 0, len(nums)-1)
	// 第二次
	reverse(nums, 0, k-1)
	// 第三次
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int){
	for start < end{
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
