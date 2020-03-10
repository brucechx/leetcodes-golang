package _066_加一

func plusOne(digits []int) []int {
	length := len(digits)
	if length == 0 {
		return []int{1}
	}

	digits[length - 1]++
	for i:=length-1; i>0; i--{
		if digits[i] < 10{
			break
		}
		digits[i] = digits[i] - 10
		digits[i-1]++
	}
	// 处理首位加1
	if digits[0] > 9{
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}
	return digits
}
