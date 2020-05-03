package _437_路径总和_III

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil{
		return 0
	}
	return helper(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func helper(root *TreeNode, sum int) int{
	if root == nil{
		return 0
	}
	sum -= root.Val
	if sum == 0{
		return 1 + helper(root.Left, sum) + helper(root.Right, sum)
	}
	return helper(root.Left, sum) + helper(root.Right, sum)
}
