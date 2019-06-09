package _59_对称的二叉树

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

/*
			 1
          2     2
        3  4  4   3
*/
func initBinaryTree() *binaryTree{
	node3L := &BinTreeNode{data:3}
	node3R := &BinTreeNode{data:3}
	node4L := &BinTreeNode{data:4}
	node4R := &BinTreeNode{data:4}
	node2L := &BinTreeNode{data:2, lChild:node3L, rChild:node4L}
	node2R := &BinTreeNode{data:2, lChild:node4R, rChild:node3R}
	root := &BinTreeNode{data:1, lChild:node2L, rChild:node2R}
	return NewBinaryTree(root)
}

func TestBinaryTree_isSymmetrical(t *testing.T){
	bt := initBinaryTree()
	assert.Equal(t, true, bt.isSymmetrical())
}
