package _151_翻转字符串里的单词

import (
	"regexp"
	"strings"
)

func reverseWords(s string) string {
	// 不用内置函数实现，字符串拼接，速度还不如内置函数快（虽然用了正则）
	// 试试使用 slice 保存， 最后拼接字符串
	res := make([]string, 0)

	left, right := len(s)-1, len(s)

	for ; left >= 0; left -- {
		if s[left] == ' ' {
			if left+1 < right {
				res = append(res, s[left+1:right])
			}

			right = left
		}
	}

	if left+1 < right {
		res = append(res, s[left+1:right])
	}


	return strings.Join(res, " ")
}


func reverseWords2(s string) string {
	// 先将字符串去除首尾空白， 然后通过空白符拆分成数组
	arr := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(s), -1)

	// 将数组反转
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	// 将数组拼接成字符串返回
	return strings.Join(arr, " ")
}