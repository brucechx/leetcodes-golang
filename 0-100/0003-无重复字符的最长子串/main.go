package _003_longest_substring_without_repeating_characters

// 使用滑动窗口[i, j),
// 如果s[j] 在 [i, j) 范围内有与 j' 重复的字符，我们不需要逐渐增加 i 。
// 我们可以直接跳过 [i，j']范围内的所有元素，并将 i 变为 j'+ 1。
func LengthOfLongestSubstring(s string) int{
	res := 0
	i := 0
	// 表示一个时间窗口
	charMap := make(map[byte]int)

	for j:=0; j< len(s); j++{
		// 如果有重复的j', i 直接变为 j' + 1
		if jj, ok := charMap[s[j]]; ok {
			if jj > i{
				i = jj
			}

		}

		currLen := j - i + 1
		if currLen > res{
			res = currLen
		}

		charMap[s[j]] = j + 1
	}
	return res
}
