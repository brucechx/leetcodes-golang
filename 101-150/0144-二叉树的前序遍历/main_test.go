package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreorderTraversal(t *testing.T){
	cases := []struct{
		root *TreeNode
		res []int
	}{
		{
			root:mockTree(),
			res: []int{1, 2, 3},
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, preorderTraversal(cas.root))
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
