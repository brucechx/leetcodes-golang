package _111_二叉树的最小深度

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDepth(t *testing.T){
	cases := []struct{
		root *TreeNode
		res int
	}{
		{
			root: mockTree1(),
			res: 2,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, minDepth(cas.root))
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
