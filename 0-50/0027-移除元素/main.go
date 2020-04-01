package _027_移除元素

// 原地置换
func removeElement(nums []int, target int) int{
	i := 0
	for _, num := range nums{
		if num != val{
			nums[i] = num
			i++
		}
	}
	return i
}

// 方法二
// 当删除元素较少的时候，直接与数组末尾交换
func removeElement2(nums []int, target int) int{
	if len(nums) == 0{
		return 0
	}
	i := 0
	n := len(nums)
	for i < n {
		if nums[i] == target{
			nums[i] = nums[n - 1]
			n --
		}else {
			i ++
		}
	}
	return n
}
