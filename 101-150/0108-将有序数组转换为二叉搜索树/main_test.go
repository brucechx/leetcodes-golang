package _108_将有序数组转换为二叉搜索树

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedArrayToBST(t *testing.T){
	cases := []struct{
		inOrder []int
		root *TreeNode
	}{
		{
			inOrder: []int{-10,-3,0,5,9},
			root: mockTree1(),
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.root, sortedArrayToBST(cas.inOrder))
	}
}



func mockTree1() *TreeNode {
	root := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   -3,
		Left:	newTreeNode(-10),
	}
	root.Right = &TreeNode{
		Val:   9,
		Left:  newTreeNode(5),
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
