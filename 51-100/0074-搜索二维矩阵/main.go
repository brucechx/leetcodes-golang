package _074_搜索二维矩阵

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0{
		return false
	}
	if target < matrix[0][0] || target > matrix[m-1][n-1]{
		return false
	}

	// get r
	r := 0
	for r < m {
		if matrix[r][0] > target{
			break
		}
		r++
	}
	r-- // 在这一行
	// 二分查找
	left, right := 0, n-1
	for left <= right{
		mid := left + (right - left) / 2
		if matrix[r][mid] == target{
			return true
		}else if matrix[r][mid] < target{
			left = mid + 1
		}else if matrix[r][mid] > target{
			right = mid - 1
		}
	}
	return false
}
