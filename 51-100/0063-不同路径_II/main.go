package _063_不同路径_II

// 动态规划
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	r := len(obstacleGrid)
	if r == 0 {
		return 0
	}
	c := len(obstacleGrid[0])
	if c == 0{
		return 0
	}
	dp := make([][]int, r)
	for i:=0; i<r; i++{
		dp[i] = make([]int, c)
	}
	if obstacleGrid[0][0] == 0{
		dp[0][0] = 1
	}

	// 行列边界
	for i:=1; i<r; i++{
		if obstacleGrid[i][0] == 0 { // 边界；第0列，如果第i行为1，则设为0；
			dp[i][0] = dp[i-1][0]
		}
	}
	for j:=1; j<c; j++{
		if obstacleGrid[0][j] == 0{
			dp[0][j] = dp[0][j-1]
		}
	}
	// 状态转移
	for i:=1;i<r;i++{
		for j:=1; j<c; j++{
			if obstacleGrid[i][j] == 0{
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[r-1][c-1]
}

