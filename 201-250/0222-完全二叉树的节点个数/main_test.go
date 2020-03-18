package _222_完全二叉树的节点个数

import (
	"fmt"
	"testing"
)

func TestCountNodes3(t *testing.T){
	fmt.Println(countNodes3(mockTree()))
}


func mockTree() *TreeNode{
	return &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val:4},
			Right: &TreeNode{Val:5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val:6},
			Right: nil,
		},
	}
}