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

func lengthOfLongestSubstring(s string) int {
	left, right, res := 0, 0, 0
	window := make(map[byte]int)
	for right < len(s){ // 窗口移动直到结束
		window[s[right]] ++
		right++
		if ! isValidWindow(window){
			tmp := s[left]
			window[tmp] --
			left ++
		}
		res = max(res, right - left)
	}
	return res
}

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}
func isValidWindow(window map[byte]int) bool{
	for _, v := range window{
		if v > 1{
			return false
		}
	}
	return true
}