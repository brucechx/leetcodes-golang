package main

// 方法一： 暴力破解
func Solution1(nums []int, target int) []int{
	for i:=0; i < len(nums); i++{
		for j:=i + 1; j < len(nums); j++{
			if nums[i] + nums[j] == target{
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 方法二： 两遍哈希
func Solution2(nums []int, target int) []int{
	numsMap := make(map[int]int)
	for i, val := range nums{
		numsMap[val] = i
	}

	for i, val := range nums{
		component := target - val
		if index, ok := numsMap[component]; ok {
			if index != i{
				return []int{i, index}
			}
		}
	}
	return []int{}
}

// 方法三： 一遍哈希
func Solution(nums []int, target int) []int{
	tmp := make(map[int]int)
	for i, v := range nums{
		t := target - v
		if _, ok := tmp[t];ok{
			return []int{tmp[t], i}
		}
		tmp[v] = i
	}
	return []int{}
}
