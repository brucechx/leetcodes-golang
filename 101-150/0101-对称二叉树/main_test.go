package _101_对称二叉树

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSymmetric(t *testing.T){
	cases := []struct{
		root *TreeNode
		res bool
	}{
		{
			root: mockTree1(),
			res: true,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, isSymmetric(cas.root))
	}
}


func mockTree1() *TreeNode {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   2,
		Left:  newTreeNode(3),
		Right: newTreeNode(4),
	}
	root.Right = &TreeNode{
		Val:   2,
		Left:  newTreeNode(4),
		Right: newTreeNode(3),
	}
	return root
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
}