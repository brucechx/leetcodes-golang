package _106_从中序与后序遍历序列构造二叉树

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T){
	cases := []struct{
		inOrder []int
		postOrder []int
		root *TreeNode
	}{
		{
			inOrder:[]int{9,3,15,20,7},
			postOrder: []int{9,15,7,20,3},
			root:mockTree1(),
		},
	}

	for _, cas := range cases{
		assert.Equal(t, buildTree(cas.inOrder, cas.postOrder), cas.root)
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
