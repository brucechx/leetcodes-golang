package _120_三角形最小路径和

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n==0{
		return 0
	}
	for i:=1; i<n; i++{
		for j:=0; j<=i; j++{
			switch  {
			case j==0:
				triangle[i][j] += triangle[i-1][0]
			case j==i:
				triangle[i][j] += triangle[i-1][j-1]
			case triangle[i-1][j-1] < triangle[i-1][j]:
				triangle[i][j] += triangle[i-1][j-1]
			default:
				triangle[i][j] += triangle[i-1][j]
			}
		}
	}

	// 从最后一行得到最后值
	min := triangle[n-1][0]
	for j:=1; j<n; j++{
		if min > triangle[n-1][j]{
			min = triangle[n-1][j]
		}
	}
	return min
}
