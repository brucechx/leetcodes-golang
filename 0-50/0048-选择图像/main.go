package _048_选择图像

import "fmt"

/*
	1。 先转置矩阵
	2。	然后每行反转
*/

/*
	123     147     741
	456 ==> 258 ==> 852
	789     369     963
 */
func rotate(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])

	// 转置矩阵； 行列互换
	for r:=0; r<row; r++{
		for c:=r; c<col; c++{
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}
	fmt.Println(matrix)
	// 每行反转
	for r:=0; r<row; r++{
		for c:=0; c<col/2; c++{
			tmp := matrix[r][c]
			matrix[r][c] = matrix[r][col-c-1]
			matrix[r][col-c-1] = tmp
		}
	}
	fmt.Println(matrix)
}