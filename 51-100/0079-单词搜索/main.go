package _079_单词搜索

// 深度搜索 + 回溯

type Solution struct {
	row int
	col int
	direction [][]int
	mark [][]bool
}

func exist(board [][]byte, word string) bool {
	row := len(board) // 行
	if row == 0{
		return false
	}
	s := NewSolution(board)

	// begin
	for r:=0; r<s.row; r++{
		for j:=0; j<s.col; j++{
			// 深度遍历
			if s.dfs(board, word, r, j, 0){
				return true
			}
		}
	}
	return false
}

func NewSolution(board [][]byte) *Solution{
	row := len(board) // 行
	col := len(board[0]) // 列
	direction := [][]int{
		{0, -1},  // 向上
		{-1, 0}, // 向左
		{0, 1}, // 向下
		{1,0}, // 向右

	}
	mark := make([][]bool, row)
	for r:=0; r<row; r++{
		mark[r] = make([]bool, col) // 默认false
	}
	return &Solution{
		row:       row,
		col:       col,
		direction: direction,
		mark:      mark,
	}
}

/*
	x,y: board 坐标
	index: word 坐标
	mark: 标注board坐标是否遍历过
 */
func (s *Solution)dfs(board [][]byte, word string, x, y,index int) bool{
	// 首先判断终止
	if index == len(word) - 1{
		return board[x][y] == word[index]
	}

	if board[x][y] == word[index]{
		s.mark[x][y] = true
		// 处理
		for _, direct := range s.direction{
			newX := x + direct[0]
			newY := y + direct[1]
			if newX <0 || newX >= s.row{
				continue
			}
			if newY <0 || newY >= s.col{
				continue
			}
			if s.mark[newX][newY]{
				continue
			}
			if s.dfs(board, word,  newX, newY, index+1){
				return true
			}
		}
		s.mark[x][y] = false
	}
	return false
}
