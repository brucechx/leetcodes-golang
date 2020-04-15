package _093_复原IP地址

import (
	"strconv"
	"strings"
)

/*
回溯法
1. 1位
2. 2位
3. 3位
*/

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	path := make([]string, 0)
	dfs(s, 0, 0, path, &res)
	return res
}

func dfs(s string, start int, count int, path []string, res *[]string){
	// 非法; 4次最大分完，还有剩余
	if len(s) - start - 3 * (4 - count) > 0{
		return
	}
	if start == len(s) && count == 4{
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, strings.Join(tmp, "."))
		return
	}
	// 大于字符长度，且分配次数大于4次
	if start >= len(s) || count == 4{
		return
	}

	// 加入一位
	path = append(path, string(s[start]))
	dfs(s, start+1, count+1, path, res)
	// 回溯
	path = path[:len(path)-1]

	// 以零开头都结束
	if s[start] == '0'{
		return
	}

	// 加入两位
	if start + 1 < len(s){
		path = append(path, s[start:start+2])
		dfs(s, start+2, count+1, path, res)
		// 回溯
		path = path[:len(path)-1]
	}

	// 加入三位
	if start + 2 < len(s){
		numS := s[start:start+3]
		num, _ := strconv.Atoi(numS)
		if num >=0 && num <=255{
			path = append(path, numS)
			dfs(s, start+3, count+1, path, res)
			// 回溯
			path = path[:len(path)-1]
		}
	}
}
