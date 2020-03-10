package _103_二叉树的锯齿形层次遍历

import (
	"fmt"
	"testing"
)

func TestPrintTree(t *testing.T){
	fmt.Println(PrintTreeNew(mockTree()))
}

func mockTree() *TreeNode{
	root := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	root.Right = &TreeNode{
		Val:   2,
		Left:  &TreeNode{
			Val:   5,
			Left:  nil,
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
	}
	return root
}
