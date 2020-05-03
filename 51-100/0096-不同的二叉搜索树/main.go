package main

/*
	G(n) = f(1)+f(2)+...+f(n)
	// 卡特兰公式
	G(n) = G(0)G(n-1) + G(1)G(n-2)+ ...+ G(n-1)G(0)
	G(n) = G(1-1)G(n-1) + G(2-1)G(n-2)+...+G(n-1)G(n-n)
 */
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i:=2; i< n+1; i++{
		for j:=1; j< i+1; j++{
			dp[i] += dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}
