package _440_字典序的第K小数字

/*
				0
			/       \
           1         2
         /  \
        10   11

 */

/*
1. 当前前缀节点数
2. k是否在当前前缀节点中
3. k不在当前前缀节点，前缀点右移
 */
func findKthNumber(n int, k int) int {
	curr := 1 // 当前游标
	prefix := 1 // 当前前缀树的根节点
	for curr < k {
		count := getCount(prefix, n)
		if curr + count > k { // 在当前前缀树中
			prefix *= 10 // 下移
			curr++ // 只会走一步
		}else { // 不在当前前缀树中
			prefix ++ // 同层；右移动；比如1 到 2
			curr += count // 走前一个前缀树中所有点； 比如走完1；10...19
		}
	}
	return prefix
}

func getCount(prefix, n int) int{
	count := 0
	curr := prefix // 比如 1
	next := prefix + 1 // 比如 2
	for curr <= n{
		count += min(next, n+1) - curr // next 大于n 的时候，非满十叉树； n +1 而非n 是因为根算一个
		curr *= 10 // 下一层；10
		next *= 10 // 20
	}
	return count
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}