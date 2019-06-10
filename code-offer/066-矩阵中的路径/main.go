package _66_矩阵中的路径

func HasPath(matrix []rune, rows, cols int, strs string) bool{
	// 边界
	if len(matrix) != rows * cols || len(strs) < 1{
		return false
	}
	var str []rune
	for _, v := range strs{
		str = append(str, v)
	}
	// 变量初始化
	visited := make([]bool, rows * cols)

	// 记录结果的数组
	pathLength := []int{0}
	// 以每一个点为起始进行搜索
	for i:=0; i< rows; i++{
		for j:=0; j<cols; j++{
			if hasPathCore(matrix, str, rows, cols, i, j, pathLength, visited){
				return true
			}
		}
	}
	return false
}

/*
	@param matrix 输入矩阵
	@param rows 矩阵行数
	@param cols 矩阵列数
	@param str 要搜索的字符
	@param visited 访问标记数组
	@param row 当前处理的行号
	@param col 当前处理的列号
	@param pathLength 已经处理的str中字符个数
 */
func hasPathCore(matrix, str []rune, rows, cols, row, col int, pathLength []int, visited []bool) bool{
	if pathLength[0] == len(str){
		return true
	}

	hasPath := false
	// 判断位置是否合法
	if row >=0 && row < rows &&
		col >= 0 && col < cols &&
		matrix[row * cols + col] == str[pathLength[0]] &&
		!visited[row * cols + col]{
		visited[row * cols + col] = true
		pathLength[0]++
		// 按左上右下进行回溯
		hasPath = hasPathCore(matrix, str, rows, cols, row, col - 1, pathLength, visited) ||
			hasPathCore(matrix, str, rows, cols, row - 1, col, pathLength, visited) ||
			hasPathCore(matrix, str, rows, cols, row, col + 1, pathLength, visited) ||
			hasPathCore(matrix, str, rows, cols, row + 1, col, pathLength, visited)
		if !hasPath {
			pathLength[0]--
			visited[row * cols + col] = false
		}

	}
	return hasPath
}
