package _052_N皇后_II

import "fmt"

/*
回溯法：
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return

    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
*/

func totalNQueens(n int) int {
	board := makeBaseBord(n)
	res := make([][]string, 0)
	backTrack(board, &res, 0)
	fmt.Println("res=", res)
	return len(res)
}

func makeBaseBord(n int) [][]byte{
	board := make([][]byte, n)
	for i:=0; i< n; i++{
		board[i] = make([]byte, n)
		for j:=0; j< n; j++{
			board[i][j] = '.'
		}
	}
	return board
}

func backTrack(board [][]byte, res *[][]string, row int){
	if row == len(board){
		boardString := make([]string, 0)
		for _, v := range board{
			boardString = append(boardString, string(v))
		}
		*res = append(*res, boardString)
		return
	}

	for col:=0; col<len(board);col++{
		if ! isValid(board, row, col){
			continue
		}
		// 做选择
		board[row][col] = 'Q'
		// 进行下一次决策
		backTrack(board, res, row+1)
		// 撤销选择, 回溯
		board[row][col] = '.'
	}
}

// 是否可以在board[row][col] 放置皇后
func isValid(board [][]byte, row, col int) bool{
	// 这一列上是否有皇后
	n := len(board)
	for i:=0; i< n; i++{
		if board[i][col] == 'Q'{
			return false
		}
	}
	// 因为按照行递增摆放,检查已摆过的行是否冲突
	// 检查右上方是否有皇后互相冲突
	rowN, colN := row - 1, col + 1
	for rowN >= 0 && colN < n{
		if board[rowN][colN] == 'Q'{
			return false
		}
		rowN--
		colN++
	}
	// 检查左上方是否有皇后互相冲突
	rowN, colN = row - 1, col - 1
	for rowN >= 0 && colN >= 0 {
		if board[rowN][colN] == 'Q'{
			return false
		}
		rowN--
		colN--
	}
	return true
}


