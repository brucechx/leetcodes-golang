package _221_最大正方形

/*
1. 如果该位置的值是 0，则 dp(i, j) = 0
2. 如果该位置的值是 1, 则 dp(i,j)=min(dp(i−1,j),dp(i−1,j−1),dp(i,j−1))+1
*/

func maximalSquare(matrix [][]byte) int {
	r := len(matrix)
	if r == 0{
		return 0
	}
	c := len(matrix[0])
	dp := make([][]int, r)
	for i:=0; i<r; i++{
		dp[i] = make([]int, c)
	}

	maxSide := 0
	for i:=0; i<r; i++{
		for j:=0; j<c; j++{
			if matrix[i][j] != '1'{
				continue
			}
			if i==0 || j == 0{
				dp[i][j] = 1
			}else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			maxSide = max(dp[i][j], maxSide)
		}
	}
	return maxSide * maxSide
}

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}
