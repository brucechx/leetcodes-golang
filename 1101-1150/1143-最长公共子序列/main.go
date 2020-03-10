package _143_最长公共子序列


func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	if m == 0 || n == 0{
		return 0
	}
	dp := make([][]int, m + 1)
	for i:=0; i<=m ; i++{
		dp[i] = make([]int, n + 1)
	}
	for i:=1; i<=m; i++{
		for j:=1; j<=n; j++{
			if text1[i-1] == text2[j-1]{
				dp[i][j] = dp[i-1][j-1] + 1
			}else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}


func longestCommonSubsequence2(text1 string, text2 string) int {
	return dp(text1, text2)
}

func dp(t1, t2 string) int{
	if t1 == "" || t2 == ""{
		return 0
	}
	i, j := len(t1)-1, len(t2)-1
	if t1[i] == t2[j]{
		return dp(t1[:i], t2[:j]) + 1
	}else {
		return max(dp(t1, t2[:j]), dp(t1[:i], t2))
	}
}

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}
