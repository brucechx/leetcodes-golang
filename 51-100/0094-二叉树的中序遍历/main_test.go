package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInorderTraversalWithRec(t *testing.T) {
	root := initTree()
	res :=InorderTraversalWithRec(root)
	expectRes := []int{3, 1, 4, 0, 5, 2}
	assert.Equal(t, expectRes, res)
}

func TestInorderTraversalWithStack(t *testing.T) {
	root := initTree()
	res := InorderTraversalWithStack(root)
	expectRes := []int{3, 1, 4, 0, 5, 2}
	assert.Equal(t, expectRes, res)
}

func initTree() *TreeNode{
	/*
	        0
	    1        2
	  3   4    5

	 */
    threeNode := &TreeNode{
    	Val:3,
    	Left:nil,
    	Right:nil,
	}
	fourNode := &TreeNode{
		Val:4,
	}
	fiveNode := &TreeNode{
		Val:5,
	}
	twoNode := &TreeNode{
		Val:2,
		Left:fiveNode,
	}
	oneNode := &TreeNode{
		Val:1,
		Left:threeNode,
		Right:fourNode,
	}
	root := &TreeNode{
		Val:0,
		Left:oneNode,
		Right:twoNode,
	}

	return root
}
