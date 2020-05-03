package _172_阶乘后的零

// n 中有多少个5
func trailingZeroes(n int) int {
	res := 0
	for n > 0 {
		res += n / 5
		n = n / 5
	}
	return res
}
