package _190_颠倒二进制位

/*
时间复杂度：O(1)，没有使用循环。
空间复杂度：O(1)，没有使用变量。
 */
func reverseBits2(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}