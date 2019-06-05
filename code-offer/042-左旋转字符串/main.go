package _42_左旋转字符串

import "fmt"

func LeftRotateString(str string, n int) (result string) {
	length := len(str)
	n = n % length
	return fmt.Sprintf("%s%s", str[n:], str[:n])
}

func LeftRotateString2(str string, n int) (result string) {
	runes := []rune(str)
	length := len(runes)
	n = n % length
	for n > 0 {
		n--
		tmp := runes[0]
		for i := 1; i < len(runes); i++ {
			runes[i-1] = runes[i]
		}
		runes[length-1] = tmp
	}
	return string(runes)
}

func LeftRotateString3(str string, n int) (result string) {
	runes := []rune(str)
	length := len(runes)
	n = n % length
	if n == 0 {
		return string(runes)
	}
	tmp := make([]rune, n)
	copy(tmp, runes)
	/*for i := n; i < length; i++ {
		runes[i-n] = runes[i]
	}*/
	copy(runes[:length-n], runes[n:])
	copy(runes[length-n:], tmp)
	return string(runes)
}

