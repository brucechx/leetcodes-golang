package v2

import "fmt"

func Fib(n int) int{
	if n == 0 || n == 1{
		return n
	}
	tmp1, tmp2, fibN := 0, 1, 0
	for i:=2;i<=n;i++{
		fibN = tmp1 + tmp2
		tmp1 = tmp2
		tmp2 = fibN
	}
	return fibN
}

func Fib2(n int) int{
	if n == 0 || n == 1{
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i:=2; i<=n; i++{
		dp[i] = dp[i-1]+dp[i-2]
	}
	fmt.Println("dp=", dp)
	return dp[n]
}
