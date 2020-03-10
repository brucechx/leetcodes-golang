package _118_杨辉三角

func generate(numRows int) [][]int {
	if numRows == 0{
		return [][]int{}
	}
	res := make([][]int, 0)
	res = append(res, []int{1})
	if numRows == 1{
		return res
	}
	for level := 1; level < numRows; level++{
		res = append(res, getNextLevel(res[level-1]))
	}
	return res
}

func getNextLevel(levelNums []int) []int{
	currentLevel := make([]int, 1, len(levelNums)+1)
	currentLevel = append(currentLevel, levelNums...)
	for i:=0; i<len(currentLevel) - 1; i++{
		currentLevel[i] += currentLevel[i+1]
	}
	return currentLevel
}

func generate2(numRows int) [][]int {
	res := make([][]int, numRows)
	for i:=0; i<numRows; i++{
		res[i] = make([]int, 0)
	}
	if numRows == 0{
		return [][]int{}
	}

	res[0] = []int{1}
	for level:=1; level<numRows; level++{
		lastLevelNums := res[level-1]
		res[level] = append(res[level], 1)
		for i:=0;i<len(lastLevelNums)-1; i++{
			tmp:= lastLevelNums[i]+lastLevelNums[i+1]
			res[level] = append(res[level], tmp)
		}
		res[level] = append(res[level], 1)
	}
	return res
}
