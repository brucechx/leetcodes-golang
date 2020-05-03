package _166_分数到小数

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0{
		return "0"
	}
	var res []string
	// 1. 结果为负数
	isN := 0
	if numerator < 0 {
		isN++
		numerator = -numerator
	}
	if denominator < 0{
		isN++
		denominator = -denominator
	}
	if isN  == 1{ // 结果为负数
		res = append(res, "-")
	}

	res = append(res, strconv.Itoa(numerator/denominator))
	reminder := numerator % denominator
	// 刚好整除
	if reminder == 0 {
		return strings.Join(res, "")
	}
	res = append(res, ".")
	m := make(map[int]int)
	for reminder != 0 {
		if v, ok := m[reminder]; ok{
			newRes := make([]string, len(res)+2)
			copy(newRes[:v], res[:v])
			newRes[v] = "("
			newRes = append(newRes, res[v:]...)
			newRes = append(newRes, ")")
			return strings.Join(newRes, "")
		}
		m[reminder] = len(res)
		reminder *= 10
		res = append(res, strconv.Itoa(reminder / denominator))
		reminder = reminder % denominator
	}
	return strings.Join(res, "")
}
