package _31_连续子数组的最大和

import "fmt"

// 递归查找
func FindGreatestSumOfSubArray(array []int) int{
	if len(array) == 1{
		return array[0]
	}

	max := array[0]
	sum := 0

	for _, v := range array{
		sum += v
		if sum > max{
			max = sum
		}
	}
	subMax := FindGreatestSumOfSubArray(array[1:])
	if max > subMax{
		return max
	}else {
		return subMax
	}
}

// 循环数组，当和小于等于0时，就把和重置成当前元素，否则就加上当前元素，判断当前和是否大于上次记录的和，大于就赋值。
func FindGreatestSumOfSubArray2(array []int) int{
	if len(array) == 1{
		return array[0]
	}

	sum := 0
	subMax := array[0]
	for _, v := range array{
		if sum <= 0{
			sum = v
		}else {
			sum += v
		}

		if sum > subMax{
			subMax = sum
		}
	}
	return subMax
}



// 循环数组，当和小于等于0时，就把和重置成当前元素，否则就加上当前元素，判断当前和是否大于上次记录的和，大于就赋值。
func FindGreatestSumOfSubArrayWithAxis(array []int) int{
	if len(array) == 1{
		return array[0]
	}

	sum := 0
	subMax := array[0]
	var axisList [][]int
	var begin, end int
	for i, v := range array{
		if sum <= 0{
			sum = v
			begin = i
		}else {
			sum += v
			end = i - 1
		}

		if sum > subMax{
			subMax = sum
			end = i
		}

		if end >= begin{
			axisList = append(axisList, []int{begin, end})
		}
		//fmt.Println("begin=", begin, " end=", end)
	}
loop:
	for _, v := range axisList{
		s := 0
		for i:=v[0]; i<=v[1]; i++{
			s += array[i]
			if s == subMax{
				fmt.Println("axis=", v)
				break loop
			}
		}
	}
	return subMax
}
