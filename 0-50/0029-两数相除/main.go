package _029_两数相除

import "math"

/*
// 比如11 / 3
   - 11比3大；3 翻倍； > 1
   - 11比6大；6 翻倍;  > 2
   - 11比12小；<4; 11 - 6 = 5; 比3大；= 3  递归
*/

func divide(m, n int) int {
	// 防止有人把0当做除数
	if n == 0 {
		return math.MaxInt32
	}

	signM, absM := abs(m)
	signN, absN := abs(n)

	res := div(absM, absN)

	// 修改res的符号
	if signM != signN {
		res = - res
	}
	// 检查溢出
	if res < math.MinInt32 || res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}

func abs(num int) (sign, abs int) {
	if num < 0{
		return -1, -num
	}
	return 1, num
}

func div(a, b int) int{
	if a < b {
		return 0
	}
	count := 1
	tmp := b // 不断更新b
	for tmp+tmp < a {
		count += count // 翻倍
		tmp += tmp // 翻倍
	}
	return count + div(a-tmp, b)
}