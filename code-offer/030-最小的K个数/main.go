package _30_最小的K个数

func GetLeastNumbersSolution(input []int, target int) []int{
	// 选择排序
	if target >= len(input){
		return input
	}
	for i:=0; i< target; i++{
		minIndex := i
		for j:= i+1; j< len(input); j++{
			if input[j] < input[minIndex]{
				minIndex = j
			}
		}
		if minIndex != i{
			input[i], input[minIndex] = input[minIndex], input[i]
		}
	}
	return input[:target]
}

