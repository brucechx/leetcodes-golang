package _14_调整数组顺序使奇数位于偶数前面

func EvenOddArray(numbers []int) (result []int) {
	var evenNumbers []int
	var oddNumbers []int

	for _, v := range numbers {
		if v%2 == 0 {
			oddNumbers = append(oddNumbers, v)
		} else {
			evenNumbers = append(evenNumbers, v)
		}
	}

	result = append(evenNumbers, oddNumbers...)
	return
}

// 前奇后偶  双指针移动
func EvenOddArray2(numbers []int) (result []int){
	if len(numbers) <= 1{
		return numbers
	}
	low := 0
	high := len(numbers) - 1

	for low < high{
		// 遇到第一个偶数退出
		for low < high && numbers[low] & 1 == 1{
			low ++
		}
		// 遇到第一个奇数退出
		for low < high && numbers[high] & 1 == 0{
			high --
		}
		numbers[low], numbers[high] = numbers[high], numbers[low]
	}
	result = numbers
	return
}
