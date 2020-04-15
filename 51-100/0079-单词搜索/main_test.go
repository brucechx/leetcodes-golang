package _079_单词搜索

import (
	"fmt"
	"testing"
)

func TestNewSolution(t *testing.T) {
	input := [][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}
	target := "ABCCED"
	fmt.Println(exist(input, target))
}
