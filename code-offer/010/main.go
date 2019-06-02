package _10

import (
	"strconv"
	"strings"
)

func V1(n int64) int{
	BINARY := 2
	// NumberOfOne 输出输入参数二进制表示的 1 的个数
	binaryString := strconv.FormatInt(n, BINARY)
	return strings.Count(binaryString, "1")
}

// NumberOfOne 输出输入参数二进制表示的 1 的个数
// flag 只有一位是 1，其他位是 0，所以如果 i 的对应为不是 1，
// 那么 i & flag 为 0
// flag 越界后变为 0，说明遍历完了
func V2(i int64) (count int) {
	var flag int64 = 1

	for flag != 0 {
		if i&flag != 0 {
			count += 1
		}
		flag <<= 1
	}

	return
}

// 把一个整数减去1， 再和原整数做与运算，会把该整数最右边一个1变为0；
func V3(i int64) (count int){
	count = 0
	for i != 0{
		count ++
		i = (i - 1) & i
	}
	return
}
