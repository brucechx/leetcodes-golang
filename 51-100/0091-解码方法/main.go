package _091_解码方法

/*
动态规划
dp[i] 为str[0,..i]的译码总数

// 分情况讨论
1. 若str[i] == '0'
	那么 str[i-1] == '1' or '2' ==> dp[i] == dp[i-2]
	否则 return 0
2. 若str[i-1] == '1'，此时str[i] != '0'
	那么 dp[i] = dp[i-1] + dp[i-2]
3. 若str[i-1] == '2', 此时str[i] != '0'
	如果str[i] <1,6>
	那么 dp[i] = dp[i-1] + dp[i-2]
*/
func numDecodings(s string) int {
	if len(s) == 0{
		return 0
	}
	if s[0] == '0'{
		return 0
	}
	pre, curr := 1, 1 // dp[-1], dp[0] = 1, 1
	for i:=1; i<len(s); i++{
		tmp := curr // dp[i]
		if s[i] == '0'{
			if s[i-1] == '1' || s[i-1] == '2'{
				curr = pre
			}else { // > 26
				return 0
			}
		}else if s[i-1] == '1' || s[i-1] == '2' && s[i]>= '1' && s[i] <= '6'{
			curr = curr + pre // dp[i] = dp[i-1] + dp[i-2]
		}
		pre = tmp
	}
	return curr
}
