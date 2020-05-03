package _130_被围绕的区域

// 转化为，如何寻找和边界联通的 OO

/*
1. dfs 与边界上O相通的设为'#'
2. 遍历替换，得解
*/
func solve(board [][]byte)  {
	if len(board) == 0{
		return
	}
	row := len(board)
	col := len(board[0])
	// 1.
	for i:=0; i<row; i++{
		for j:=0; j<col; j++{
			// 从边缘开始搜索
			isEdge := i==0 || j==0 || i==row-1 || j == col -1
			if isEdge && board[i][j] == 'O'{
				dfs(board, i, j, row, col)
			}
		}
	}
	// 2.
	for i:=0; i<row; i++{
		for j:=0; j<col; j++{
			if board[i][j] == '#'{
				board[i][j] = 'O'
			}else if board[i][j] == 'O'{
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, i, j, row, col int){
	if i < 0 || i >= row || j<0 || j>=col || board[i][j] == 'X' || board[i][j] == '#'{
		return
	}
	board[i][j] = '#'
	dfs(board, i-1, j, row, col)
	dfs(board, i+1, j, row, col)
	dfs(board, i, j-1, row, col)
	dfs(board, i, j+1, row, col)
}