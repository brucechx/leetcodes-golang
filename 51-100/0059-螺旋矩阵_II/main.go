package _059_螺旋矩阵_II

/*
	向内旋转

*/

func generateMatrix(n int) [][]int{
	if n == 0 {
		return nil
	}
	target := n * n
	left, right, top, bottom := 0, n-1, 0, n-1
	num := 1
	matrix := initMatrix(n)
	// 不用left<right || top < bottom 是因为当n 为奇数时，中间点无法填充
	for num <= target{
		// left --> right
		for i:=left; i<=right; i++{
			matrix[top][i] = num
			num++
		}
		top++
		// top --> bottom
		for i:=top; i<=bottom; i++{
			matrix[i][right] = num
			num++
		}
		right--
		// right --> left
		for i:=right; i>=left; i--{
			matrix[bottom][i] = num
			num++
		}
		bottom--
		// bottom --> top
		for i:=bottom; i>=top; i--{
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}

func initMatrix(n int) [][]int{
	matrix := make([][]int, n)
	for i:=0; i<n; i++{
		matrix[i] = make([]int, n)
	}
	return matrix
}