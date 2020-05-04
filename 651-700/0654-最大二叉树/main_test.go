package _654_最大二叉树

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T){
	input := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(input)
	fmt.Println(root.Val)
}
