package _239_滑动窗口最大值

// 双向队列
func maxSlidingWindow(nums []int, k int) []int {
	// 边界
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	q := make([]int, 0)
	res := make([]int, 0)
	// 初始化前k个
	max := 0
	for i:=0; i<k; i++{
		q = cleanQueue(q, nums, i, k)
		q = append(q, i)
		if nums[i] > nums[max]{
			max = i
		}
	}
	res = append(res, nums[max])

	// 第k
	for i:=k; i<len(nums); i++{
		q = cleanQueue(q, nums, i, k)
		q = append(q, i)
		res = append(res, nums[q[0]])
	}
	return res
}

// clean_queue
func cleanQueue(q []int, nums []int, i, k  int) []int{
	// queue 中存放数组下标
	kIndex := i - k
	// 移除滑动窗口之外的下标
	for len(q) > 0  && q[0] <= kIndex{
		q = q[1:]
	}
	// 移除比让前值小的下标
	for len(q) > 0 && nums[i] > nums[q[len(q)-1]]{
		q = q[:len(q)-1]
	}
	return q
}
