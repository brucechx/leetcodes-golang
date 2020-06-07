package _201_数字范围按位与

//m < n，所以至少会有两个数字，所以最低位相与结果一定是 0。
func rangeBitwiseAnd(m int, n int) int {
	for m < n {
		n = n & (n - 1)
	}
	return n
}
