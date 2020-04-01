package _028_实现strStr

// KMP
// 简单直接匹配，复杂度高
func strStr(haystack string, needle string) int {
	hLen, nLen := len(haystack), len(needle)

	for i:=0; i<=hLen - nLen; i++{
		if haystack[i:i+nLen] == needle{
			return i
		}
	}
	return -1
}
