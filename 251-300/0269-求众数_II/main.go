package _269_求众数_II

// 摩尔投票
func majorityElement(nums []int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}

	cand1, cand2 := nums[0], nums[0]
	count1, count2 := 0, 0
	for _, num := range nums{
		// 如果下一个选票是候选者，则候选者当选的计数加一，反之减一
		if num == cand1{
			count1++
			continue
		}
		if num == cand2{
			count2++
			continue
		}
		// 如果当前候选的计数为0，则变更候选者
		if count1 == 0 {
			cand1 = num
			count1++
			continue
		}
		if count2 == 0{
			cand2 = num
			count2++
			continue
		}
		// 当前选票都不是候选者，则当前候选者全部减一；M个候选者，同理
		count1 --
		count2 --
	}
	// 计数阶段
	count1 = 0
	count2 = 0
	for _, num := range nums {
		if cand1 == num {
			count1++
		} else if cand2 == num {
			count2++
		}
	}
	if count1 > len(nums)/3 {
		res = append(res, cand1)
	}
	if count2 > len(nums)/3 {
		res = append(res, cand2)
	}
	return res
}
