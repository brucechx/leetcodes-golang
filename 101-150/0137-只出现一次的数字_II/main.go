package _137_只出现一次的数字_II

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums{
		ones = ones ^ num &^twos
		twos = twos ^ num &^ones
	}
	return ones
}
