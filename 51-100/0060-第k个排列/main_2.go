package _060_第k个排列

// 康托编码
/*
如何找出第16个（按字典序的）{1,2,3,4,5}的全排列？
首先用16-1得到15
用15去除4! 得到0余15
用15去除3! 得到2余3
用3去除2! 得到1余1
用1去除1! 得到1余0

有0个数比它小的数是1，所以第一位是1
有2个数比它小的数是3，但1已经在之前出现过了所以是4
有1个数比它小的数是2，但1已经在之前出现过了所以是3
有1个数比它小的数是2，但1,3,4都出现过了所以是5
最后一个数只能是2

所以排列为1 4 3 5 2
*/

func getPermutation2(n int, k int) string {
	if n == 0 {
		return ""
	}

	// 待存放的数字
	params := make([]byte, n)
	for i:=0; i<n; i++{
		params[i] = byte(i) + '1'
	}
	// 从0开始
	k--
	// n!
	base := factorial(n)

	// 存放结果
	res := make([]byte, 0)

	for i:=n-1; i>0; i--{
		index := k/base
		val := params[index]
		res = append(res, val)
		// 移除已经选过的
		params = append(params[:index], params[index+1:]...)
		k %= base
		base /= i
	}
	// 最后一个数
	res = append(res, params[0])
	return string(res)
}

func factorial(n int) int{
	result := 1
	for i:=1;i<n;i++{
		result *= i
	}
	return result
}