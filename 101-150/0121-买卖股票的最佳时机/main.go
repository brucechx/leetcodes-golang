package _121_买卖股票的最佳时机

import "math"

/*
base-case
dp[-1][k][0] = dp[i][0][0] = 0
dp[-1][k][1] = dp[i][0][1] = -infinity

==>
dp[-1][1][0] = dp[i][0][0] = 0
dp[-1][1][1] = dp[i][0][1] = -infinity

dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k-1][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])

==> k=1 对状态转移无影响
dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
dp[i][1] = max(dp[i-1][1], -prices[i])
*/

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0{
		return 0
	}
	dp := make([][]int, n)
	for i:=0; i<n; i++{
		dp[i] = make([]int, 2)
		if i == 0{
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}


func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 0{
		return 0
	}
	dp_i_0, dp_i_1 := 0, math.MinInt32
	for i:=0; i<n; i++{
		dp_i_0 = max(dp_i_0, dp_i_1 + prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}
	return dp_i_0
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}

