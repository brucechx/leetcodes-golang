package _6_数组中数字出现的次数

/*
1. 先全部异或
2. 找到区分位
3. 分组异或；得解
*/

func singleNumbers(nums []int) []int {
	k := 0
	// 1
	for _, num := range nums{
		k ^= num
	}

	// 2
	mask := 1
	for k & mask == 0 {
		mask <<= 1
	}

	// 3
	a, b := 0, 0
	for _, num := range nums{
		if num & mask == 0 { // 基于区分位，分组
			a ^= num
		}else {
			b ^= num
		}
	}
	return []int{a, b}
}
