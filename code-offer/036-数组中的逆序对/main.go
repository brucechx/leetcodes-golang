package _36_数组中的逆序对

func InversePairs(array []int) int{
	res := 0

	for i:=0; i < len(array); i++{
		for j:=i+1; j< len(array); j++{
			if array[i] > array[j]{
				res++
			}
		}
	}
	return res
}
