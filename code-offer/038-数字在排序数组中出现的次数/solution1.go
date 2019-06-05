package _38_数字在排序数组中出现的次数

// 暴力法
// 遍历1遍数组
func GetNumberOfK(data []int, k int) (result int) {
	for _, v := range data {
		if k == v {
			result++
		}
	}
	return
}
