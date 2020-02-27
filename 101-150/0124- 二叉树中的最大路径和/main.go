package _124__二叉树中的最大路径和

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil{
		return 0
	}
	maxSum := root.Val
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil{
			return 0
		}
		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))
		sum := left + root.Val + right
		if sum > maxSum{
			maxSum = sum
		}
		return max(left, right) + root.Val
	}
	dfs(root)
	return maxSum
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}
