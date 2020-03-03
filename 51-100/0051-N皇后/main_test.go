package _051_N皇后

import (
	"fmt"
	"testing"
)

func TestSolveNQueens(t *testing.T){
	res := solveNQueens(4)
	fmt.Println(res)
	fmt.Println(len(res))
}
