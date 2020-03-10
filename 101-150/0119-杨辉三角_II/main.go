package _119_杨辉三角_II

func getRow(rowIndex int) []int {
	rowIndex++
	if rowIndex == 0{
		return []int{}
	}
	res := make([][]int, 0)
	res = append(res, []int{1})
	if rowIndex == 1{
		return []int{1}
	}
	for level:=1; level<=rowIndex; level++{
		res = append(res, getNext(res[level-1]))
	}
	return res[rowIndex]
}

func getNext(levelNums []int) []int{
	res := make([]int, 1, len(levelNums)+1)
	res = append(res, levelNums...)
	for i:=0; i<len(res)-1; i++{
		res[i] += res[i+1]
	}
	return res
}
