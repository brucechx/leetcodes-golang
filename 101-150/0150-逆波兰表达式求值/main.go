package _150_逆波兰表达式求值

import "strconv"

func evalRPN(tokens []string) int {
	// 保存数字的栈
	nums := make([]int, 0, len(tokens))
	for _, token := range tokens{
		switch token {
		case "+", "-", "*", "/": // 运算符
			b, a := nums[len(nums)-1], nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			nums = append(nums, compute(a, b, token))
		default: // 数字
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	return nums[0]
}

func compute(a, b int, s string) int{
	switch s {
	case "+":
		return a+b
	case "-":
		return a-b
	case "*":
		return a*b
	case "/":
		return a/b
	}
	return -1
}
