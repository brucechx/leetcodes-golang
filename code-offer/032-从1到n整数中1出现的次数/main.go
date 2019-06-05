package _32_从1到n整数中1出现的次数

func NumberOf1Between1AndN_Solution1(n int) (res int) {
	for i := 1; i <= n; i++ {
		num := i
		for {
			if num == 0 {
				break
			}
			if num%10 == 1 {
				res++
			}
			num = num / 10
		}
	}
	return res
}

/*
总结一下以上的算法，可以看到，当计算右数第 i 位包含的 X 的个数时：
取第 i 位左边（高位）的数字，乘以 10i−1，得到基础值 a。
取第 i 位数字，计算修正值：
如果大于 X，则结果为 a+10i−1。
如果小于 X，则结果为 a。
如果等 X，则取第 i 位右边（低位）数字，设为 b，最后结果为 a+b+1。
*/
func NumberOf1Between1AndNSolution2(n int) int {
	count := 0
	for i := 1; i <= n; i *= 10 {
		a := n / i
		b := n % i
		//之所以补8，是因为当百位为0，则a/10==(a+8)/10，
		//当百位>=2，补8会产生进位位，效果等同于(a/10+1)
		count += (a + 8) / 10 * i
		if a%10 == 1 {
			// 如果是1xx...,就会有 b+1 个1
			count += b + 1
		}
	}
	return count
}
