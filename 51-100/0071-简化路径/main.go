package _071_简化路径

import "strings"

// 用栈模拟目录进退

func simplifyPath(path string) string {
	params := strings.Split(path, "/")
	stack := make([]string, 0)
	for _, param := range params{
		// 为空或者. 时忽略
		if param == "" || param == "."{
			continue
		}
		// 为..时，如果栈中有数据则pop栈顶
		if param == ".."{
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}else {
			stack = append(stack, param)
		}
	}
	return "/"+strings.Join(stack, "/")
}
