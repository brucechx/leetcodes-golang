package _226_翻转二叉树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil){
		return root
	}
	root.Left = invertTree(root.Right)
	root.Right = invertTree(root.Left)
	return root
}
