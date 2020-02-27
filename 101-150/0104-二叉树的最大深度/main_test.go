package _104_二叉树的最大深度

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T){
	cases := []struct{
		root *TreeNode
		depth int
	}{
		{
			root:mockTree1(),
			depth: 3,
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.depth, maxDepth(cas.root))
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
