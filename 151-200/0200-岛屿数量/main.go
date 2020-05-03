package _200_岛屿数量

// dfs
func numIslands(grid [][]byte) int {
	res := 0
	row := len(grid)
	if row == 0{
		return 0
	}
	col := len(grid[0])
	for r:=0; r<row; r++{
		for c:=0; c<col; c++{
			if grid[r][c] == '1'{
				res++
				dfs(grid, r, c, row, col)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, r, c int, row, col int){
	grid[r][c] = '0'
	if r-1 >=0 && grid[r-1][c] == '1'{
		dfs(grid, r-1, c, row, col)
	}
	if r+1 < row && grid[r+1][c] == '1'{
		dfs(grid, r+1, c, row, col)
	}
	if c-1 >=0 && grid[r][c-1] == '1'{
		dfs(grid, r, c-1, row, col)
	}
	if c+1 < col && grid[r][c+1] == '1'{
		dfs(grid, r, c+1, row, col)
	}

}
