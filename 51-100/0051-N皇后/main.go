package _051_N皇后

/*
     _____________>
	|
  	|
	|

 */
func solveNQueens(n int) [][]string {
	board := makeBoard(n)
	var res [][]string
	solveNQueen(board, 0, &res)
	return res
}

func solveNQueen(board []string, row int, res *[][]string){
	// 触发结束条件
	if row == len(board){
		tmp := make([]string, len(board))
		copy(tmp, board)
		*res = append(*res, tmp)
		return
	}
	n := len(board)
	for col:=0; col<n; col++{
		// 排除不合法选择
		if !isValid(board, row, col){
			continue
		}
		// 做选择
		board[row] = markBoard(board[row], col, 'Q')
		// 进入下一行决策
		solveNQueen(board, row + 1, res)
		// 撤销选择
		board[row] = markBoard(board[row], col, '.')
	}
}

func markBoard(input string, col int, v byte) string{
	tmp := []byte(input)
	for i:=0; i<len(tmp); i++{
		if col == i{
			tmp[i] = v
		}
	}
	return string(tmp)
}

func isValid(board []string, row, col int) bool{
	n := len(board)
	// 检查列
	for i:=0; i<n; i++{
		if board[i][col] == 'Q'{
			return false
		}
	}
	// 检查右上方是否有皇后互相冲突
	var i, j int
	i = row - 1
	j = col + 1
	for i>=0 && j<n{
		if board[i][j] == 'Q'{
			return false
		}
		i--
		j++
	}

	// 检查左上方是否有皇后互相冲突
	i = row - 1
	j = col - 1
	for i>=0 && j>=0{
		if board[i][j] == 'Q'{
			return false
		}
		i--
		j--
	}
	return true
}


func makeBoard(n int) []string{
	var boardByte []byte
	for i:=0; i<n; i++{
		boardByte = append(boardByte, '.')
	}
	boardStr := string(boardByte)
	var board []string
	for i:=0; i<n; i++{
		board = append(board, boardStr)
	}
	return board
}


