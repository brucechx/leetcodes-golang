package _54_表示数值的字符串


func IsNumeric(str string) bool{
	if str == "" || len(str) < 1{
		return false
	}

	index := 0
	if str[index] == '+' || str[index] == '-'{
		index++
	}
	// 已经到达字符串的末尾了
	if index >= len(str){
		return false
	}
	index = scanDigits(str, index)
	// 如果到达字符尾
	if index == len(str){
		return true
	}
	// 如果是小数点
	switch str[index] {
	case '.':
		// 移动到下一个位置
		index++
		index = scanDigits(str, index)
		// 已经到了字符串的末尾了
		if index >= len(str){
			return true
		}else if str[index] == 'e' || str[index] == 'E'{
			return isExponential(str, index)
		}else {
			return false
		}
	case 'E', 'e':
		return isExponential(str, index)
	}
	return false
}

// 扫描字符串部分的数字部分
func scanDigits(str string, index int) int{
	for index < len(str) && str[index] >= '0' && str[index] <= '9'{
		index++
	}
	return index
}

// 判断是否是科学计数法的结尾部分，如E5，e5，E+5，e-5，e(E)后面接整数
func isExponential(str string, index int) bool{
	if index >= len(str) || str[index] != 'E' && str[index] != 'e'{
		return false
	}
	// 移动到下一个要处理的位置
	index++
	// 到达字符串的末尾，就返回false
	if index >= len(str){
		return false
	}
	if str[index] == '+' || str[index] == '-'{
		index++
	}
	if index >= len(str){
		return false
	}
	index = scanDigits(str, index)
	// 如果已经处理到了的数字的末尾就认为是正确的指数
	return index >= len(str)
}

