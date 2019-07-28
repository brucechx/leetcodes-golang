package _67_机器人的运动范围

/*
	threshold 阀值
 */
func MovingCount(threshold, rows, cols int) int{
	// 边界
	if threshold < 0 || rows < 1 || cols < 1{
		return 0
	}
	visited := make([]bool, rows * cols)
	return movingCountCore(threshold, rows, cols, 0, 0, visited)
}

// 递归回溯
func movingCountCore(threshold, rows, cols, row, col int, visited []bool) int{
	count := 0
	if check(threshold, rows, cols, row, col, visited){
		visited[row * cols + col] = true
		count = 1+ movingCountCore(threshold, rows, cols, row-1, col, visited) +
			movingCountCore(threshold, rows, cols, row, col-1, visited) +
			movingCountCore(threshold, rows, cols, row+1, col, visited) +
			movingCountCore(threshold, rows, cols, row, col+1, visited)
	}
	return count
}

// 断机器人能否进入坐标为(row, col)的方格

func check(threshold, rows, cols, row, col int, visited []bool) bool{
	return row >=0 && row < rows &&
			col >=0 && col < cols &&
			!visited[row * cols + col] &&
			(getDigitSum(row) + getDigitSum(col) <= threshold)
}

// 一个数字的数位之和
func getDigitSum(num int) int{
	result := 0
	for num >0{
		result += num % 10
		num /= 10
	}
	return result
}