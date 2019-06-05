package _38_数字在排序数组中出现的次数

// 二分查找法
func GetNumberOfK2(data []int, target int) int{
	middle, start, end := 0, -1, -1
	left := 0
	right := len(data) - 1

	for left <= right{
		middle = (left + right)/ 2
		if data[middle] == target{
			tmp := middle
			for tmp >= left && data[tmp] == target{
				start = tmp
				tmp --
			}
			tmp = middle
			for tmp <= right && data[tmp] == target{
				end = tmp
				tmp ++
			}
			break
		}else if data[middle] > target{
			right = middle - 1
		}else {
			left = middle + 1
		}
	}
	if start == -1 || end == -1{
		return 0
	}
	return end - start + 1
}
