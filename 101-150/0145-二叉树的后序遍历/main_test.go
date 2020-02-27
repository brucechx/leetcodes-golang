package _145_二叉树的后序遍历

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostorderTraversal(t *testing.T){
	cases := []struct{
		root *TreeNode
		res []int
	}{
		{
			root:mockTree(),
			res: []int{3, 2, 1},
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, postorderTraversal(cas.root))
	}
}


func mockTree() *TreeNode {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   2,
		Left:  newTreeNode(3),
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
