package _005_最大回文数

import "fmt"

// 中心扩展算法

func longestPalindrome(s string) string {
	if s == "" || len(s) < 1{
		return ""
	}
	start := 0
	end := 0
	for i:=0; i<len(s);i++{
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		length := len1
		if len1 < len2{
			length = len2
		}
		if length > (end - start){
			start = i - (length - 1) / 2
			end = i + length / 2
		}
		fmt.Println(i, " length=", length, " start=", start, " end=", end)
	}
	return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) int{
	l := left
	r := right
	for {
		if l < 0 {
			break
		}
		if r >= len(s){
			break
		}
		if s[l] != s[r]{
			break
		}
		l --
		r ++
	}
	fmt.Println("r=", r, " l=", l)
	return r - l - 1
}