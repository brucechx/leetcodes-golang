package _139_单词拆分

/*
- dp[i]：表示前i个元素是否能拆分成词典中的单词，故最终返回dp[len(s)]。

- 状态转移：初始化dp[0]=true，对每一个起始元素i遍历截止元素j（j范围[i+1, len(s)]），若前i个元素已经可以成功拆分且i到j也是一个可选单词，
        则前j个元素也可以成功拆分，即d[j]置为true。
*/

func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return false
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i:=0; i<len(s); i++{
		for j:=i+1; j<len(s)+1; j++{
			for _, w := range wordDict{
				if s[i:j] == w && dp[i] {
					dp[j] = true
				}
			}
		}
	}
	return dp[len(s)]
}
