package _179_最大数

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	numsS := make([]string, len(nums))
	// 首先转为字符串
	for i, num := range nums{
		numsS[i] = strconv.Itoa(num)
	}
	// 定义字符串排序
	sort.Slice(numsS, func(i, j int) bool {
		return numsS[i] + numsS[j] >= numsS[j] + numsS[i]
	})
	res := strings.Join(numsS, "")
	if res[0] == '0'{
		return "0"
	}
	return res
}
