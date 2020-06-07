package _227_基本计算器II

func calculate(s string) int {
	num := 0
	sign := byte('+')
	stack := make([]int, 0)
	for i, v := range []byte(s){
		if isdigit(v){
			num = num * 10 + int(v - '0')
		}
		if ! isdigit(v) && v != ' ' || i == len(s)-1{
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				pre := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pre * num)
			case '/':
				pre := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pre / num)
			}
			sign = v
			num = 0
		}

	}
	sum := 0
	for _, ss := range stack{
		sum += ss
	}
	return sum
}

func isdigit(c byte) bool{
	return c >= '0' && c <= '9'
}
