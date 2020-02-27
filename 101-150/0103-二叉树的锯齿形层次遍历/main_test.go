package _103_二叉树的锯齿形层次遍历

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZigzagLevelOrder(t *testing.T){
	cases := []struct{
		root *TreeNode
		res [][]int
	}{
		{
			root: mockTree1(),
			res: [][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, zigzagLevelOrder(cas.root))
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