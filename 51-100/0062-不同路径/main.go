package _062_不同路径

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i:=0; i<m; i++{
		dp[i] = make([]int, n)
		dp[i][0] = 1 // 第一列，状态方程中的边界；j=0, 只有一条路线
	}

	for j:=0; j<n; j++{
		dp[0][j] = 1 // 第一行，边界，路径只有一条
	}
	for i:=1; i<m; i++{
		for j:=1; j<n; j++{
			dp[i][j] = dp[i-1][j]+dp[i][j-1] // 状态转移
		}
	}
	return dp[m-1][n-1]
}
