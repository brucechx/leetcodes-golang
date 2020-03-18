package _136_只出现一次的数字

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums{
		res ^=num
	}
	return res
}