package _020_有效的括号

func isValid(str string) bool{
	charMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	// 构建栈
	stack := make([]rune, 0, len(str))
	for _, v := range str{
		c, ok := charMap[v]
		if !ok { // 开括号，入栈
			stack = append(stack, v)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		// 闭括号，与栈顶比较，
		stackTop := stack[len(stack) - 1]
		stack = stack[: len(stack) - 1]
		if c != stackTop{
			return false
		}
	}
	return len(stack) == 0
}
