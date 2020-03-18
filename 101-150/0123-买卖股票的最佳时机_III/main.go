package _123_买卖股票的最佳时机_III

import (
	"fmt"
	"math"
)

/*
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k-1][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])

*/
func maxProfit(prices []int) int{
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][][]int, n)
	for i:=0; i<n; i++{
		dp[i] = make([][]int, 3)
		for k:=0; k<=2; k++{
			dp[i][k] = make([]int, 2)
		}
	}
	//dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1] + prices[i])
	//dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0] - prices[i])
	//dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
	//dp[i][1][1] = max(dp[i-1][1][1], -prices[i])
	// ​
	dp_i10, dp_i11 := 0, math.MinInt32
	dp_i20, dp_i21 := 0, math.MinInt32
	for i, price := range prices {
		dp_i20 = max(dp_i20, dp_i21 + price)
		dp_i21 = max(dp_i21, dp_i10 - price)
		dp_i10 = max(dp_i10, dp_i11 + price)
		dp_i11 = max(dp_i11, -price)
	}
	return dp_i20
}
func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}

