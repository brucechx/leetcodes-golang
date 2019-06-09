package _52_构建乘积数组

func Multiply(arr []int) []int{
	n := len(arr)
	res := make([]int, n)
	// c[i] = a[0]*a[1]*...*a[i-1]
	tmp := 1
	for i:=0; i<n ; i++{
		res[i] = tmp
		tmp *= arr[i]
	}
	// d[i] = a[n-1]*a[n-2]*...*a[i+1]
	tmp = 1
	for i:=n-1; i>=0; i--{
		res[i] *= tmp
		tmp *= arr[i]
	}
	return res
}
