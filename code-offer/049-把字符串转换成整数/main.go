package _49_把字符串转换成整数

import (
	"strings"
	"math"
)

func StrToInt(str string) (result int) {
	str = strings.TrimSpace(str)
	if str == "" {
		return
	}
	// 正负号
	i := 1
	if strings.HasPrefix(str, "+") {
		i = 1
		str = strings.TrimLeft(str, "+")
	} else if strings.HasPrefix(str, "-") {
		i = -1
		str = strings.TrimLeft(str, "-")
	}

	//循环每个数字
	for _, v := range []rune(str) {
		if v < '0' || v > '9'{
			return 0
		}
		// 0-9为合法

		num := int(v - '0')

		// 溢出
		if result > math.MaxInt32-num {
			return 0
		}

		result = result*10 + num

		// 溢出
		if result < 0 {
			return 0
		}
	}
	return result * i
}
