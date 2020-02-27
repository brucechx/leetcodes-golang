package _112_路径总和

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPathSum(t *testing.T){
	cases := []struct{
		root *TreeNode
		sum int
		res bool
	}{
		{
			root:mockTree(),
			sum: 22,
			res: true,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, hasPathSum(cas.root, cas.sum))
	}
}


func mockTree() *TreeNode {
	root := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   4,
		Left: newTreeNode(11),
	}
	root.Left.Left.Left = newTreeNode(7)
	root.Left.Left.Right = newTreeNode(2)
	root.Right = &TreeNode{
		Val:   8,
		Left:  newTreeNode(13),
		Right: newTreeNode(4),
	}
	root.Right.Right.Right = newTreeNode(1)
	return root
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
