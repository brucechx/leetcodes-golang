package _199_二叉树的右视图

import (
	"fmt"
	"testing"
)

func TestRightSideView(t *testing.T){
	fmt.Println(rightSideView(mockTree()))
} 

func mockTree() *TreeNode{
	return &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: &TreeNode{Val:5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
}