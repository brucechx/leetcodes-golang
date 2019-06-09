package _58_二叉树的下一个结点

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBinaryTree_NextNode(t *testing.T) {
	binaryTree := initBinaryTree()
	cases := []struct{
		node *BinTreeNode
		res int
	}{
		{binaryTree.Find(2), 4},
		{binaryTree.Find(4), 3},
		{binaryTree.Find(3), 1},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, binaryTree.InOrderNextNode(cas.node).GetData())
	}
}

func initBinaryTree() *binaryTree{
	a := NewBinTreeNode(1)
	tree := NewBinaryTree(a)
	a.SetLChild(NewBinTreeNode(2))
	a.SetRChild(NewBinTreeNode(5))
	a.GetLChild().SetRChild(NewBinTreeNode(3))
	a.GetLChild().GetRChild().SetLChild(NewBinTreeNode(4))
	a.GetRChild().SetLChild(NewBinTreeNode(6))
	a.GetRChild().SetRChild(NewBinTreeNode(7))
	a.GetRChild().GetLChild().SetRChild(NewBinTreeNode(9))
	a.GetRChild().GetRChild().SetRChild(NewBinTreeNode(8))
	return tree
}
