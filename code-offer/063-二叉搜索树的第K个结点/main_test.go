package _63_二叉搜索树的第k结点

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

/*
            4
           /  \
         2     6
        / \   / \
       1   3 5   7

// 1234567
*/

func TestBinarySearchTree_kthNode(t *testing.T) {
	binarySearchTree := initBinaryTree()
	assert.Equal(t, 2, binarySearchTree.kthNode(2).data)
	assert.Equal(t, 4, binarySearchTree.kthNode(4).data)
	assert.Equal(t, 6, binarySearchTree.kthNode(6).data)
}

func initBinaryTree() *binarySearchTree{
	 node1 := &BinTreeNode{data:1}
	 node3 := &BinTreeNode{data:3}
	 node5 := &BinTreeNode{data:5}
	 node7 := &BinTreeNode{data:7}
	 node2 := &BinTreeNode{data:2, lChild:node1, rChild:node3}
	 node6 := &BinTreeNode{data:6, lChild:node5, rChild:node7}
	 root := &BinTreeNode{data:4, lChild:node2, rChild:node6}
	 return NewBinarySearchTree(root)
}