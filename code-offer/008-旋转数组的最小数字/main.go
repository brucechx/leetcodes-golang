package _08_旋转数组的最小数字

// 基于规律，通过二分查找解决
func FindMin(num []int) int{
	length := len(num)
	if length <= 0 {
		panic("nums is null")
	}
	// for 场景一
	index1 := 0 // 没有旋转
	index2 := length - 1
	midIndex := index1

	for num[index1] >= num[index2]{
		//  如果前一个元素与后一个元素差一位
		//  说明找到了最大最小的元素
		// 场景二
		if index2 - index1 == 1{
			midIndex = index2
			break
		}

		midIndex = (index1 + index2) / 2

		// for 场景三
		if num[index1] == num[index2] && num[midIndex] == num[index2]{
			return MinInOrder(num, index1, index2)
		}
		if num[midIndex] >= num[index1]{
			index1 = midIndex
		}else if num[midIndex] <= num[index2]{
			index2 = midIndex
		}
	}
	return num[midIndex]
}

func MinInOrder(num []int, index1, index2 int) int{
	result := num[index1]
	for i:=index1;i<=index2;i++{
		if num[i] < result{
			result = num[i]
		}
	}
	return result
}