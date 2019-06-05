package _29_数组中出现次数超过一半的数字

func MoreThanHalfNumSolution(array []int) int{
	numMap := make(map[int]int)
	arraySize := len(array)
	for _, v := range array{
		if val, ok := numMap[v]; ok {
			val++
			if val >= arraySize/2{
				return v
			}
			numMap[v] = val
		}else {
			numMap[v] = 1
		}
	}
	return 0
}
