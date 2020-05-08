package _671_二叉树中第二小的节点

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil{
		return -1
	}
	return helper(root, root.Val)
}

func helper(root *TreeNode, min int) int{
	if root == nil{
		return -1
	}
	if root.Val > min{
		return root.Val
	}
	l := helper(root.Left, min)
	r := helper(root.Right, min)
	if l == -1{
		return r
	}
	if r == -1{
		return l
	}
	return minN(l, r)
}

func minN(a, b int) int{
	if a < b{
		return a
	}
	return b
}
