package _653_两数之和_IV__输入BST

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	return helper(root, root, k)
}

func helper(node *TreeNode, root *TreeNode, k int) bool{
	if node == nil{
		return false
	}
	return node.Val * 2 != k && findNode(root, k - node.Val) ||
		helper(node.Left, root, k) || helper(node.Right, root, k)
}

func findNode(root *TreeNode, target int) bool {
	if root == nil{
		return false
	}
	if root.Val == target{
		return true
	}
	if root.Val < target{
		return findNode(root.Right, target)
	}
	return findNode(root.Left, target)
}

