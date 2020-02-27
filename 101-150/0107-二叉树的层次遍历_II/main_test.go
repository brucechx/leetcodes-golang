package _107_二叉树的层次遍历_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrderBottom(t *testing.T){
	cases := []struct{
		root *TreeNode
		res [][]int
	}{
		{
			root: mockTree1(),
			res: [][]int{
				{15, 7},
				{9, 20},
				{3},
			},
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, levelOrderBottom(cas.root))
	}
}



func mockTree1() *TreeNode {
	root := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   9,
	}
	root.Right = &TreeNode{
		Val:   20,
		Left:  newTreeNode(15),
		Right: newTreeNode(7),
	}
	return root
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
