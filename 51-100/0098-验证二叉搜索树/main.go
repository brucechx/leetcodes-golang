package main

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// Go int 类型的最小值与最大值
	//MIN, MAX := -1<<63, 1<<63-1
	MIN, MAX := math.MinInt64, math.MaxInt64
	return isBST(MIN, MAX, root)
}

//
func isBST(min, max int, node *TreeNode) bool {
	if node == nil {
		return true
	}
	return min < node.Val &&
		node.Val < max &&
		isBST(min, node.Val, node.Left) &&
		isBST(node.Val, max, node.Right)
}

