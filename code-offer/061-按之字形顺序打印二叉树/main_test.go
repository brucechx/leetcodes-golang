package _61_按之字形顺序打印二叉树

import "testing"

/*
			 1
          2     3
        4  5  6   7
*/
func initBinaryTree() *binaryTree{
	node4 := &BinTreeNode{data:4}
	node5 := &BinTreeNode{data:5}
	node6 := &BinTreeNode{data:6}
	node7 := &BinTreeNode{data:7}
	node2 := &BinTreeNode{data:2, lChild:node4, rChild:node5}
	node3 := &BinTreeNode{data:3, lChild:node6, rChild:node7}
	root := &BinTreeNode{data:1, lChild:node2, rChild:node3}
	return NewBinaryTree(root)
}

func TestBinaryTree_PrintZ(t *testing.T) {
	bt := initBinaryTree()
	bt.PrintZ()
	/*
	output:
			1
			32
			4567
	*/
}
