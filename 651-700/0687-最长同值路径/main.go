package _687_最长同值路径

type TreeNode struct {
	Val int
    Left *TreeNode
	Right *TreeNode
}


func longestUnivaluePath(root *TreeNode) int {
	res := 0
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil{
			return 0
		}
		left := helper(node.Left)
		right := helper(node.Right)
		arrowLeft, arrowRight := 0, 0
		if node.Left != nil && node.Left.Val == node.Val{ // 左子等于父
			arrowLeft += left + 1
		}
		if node.Right != nil && node.Right.Val == node.Val{ // 右子等于父
			arrowRight += right + 1
		}
		res = max(res, arrowLeft+arrowRight)
		return max(arrowLeft, arrowRight)
	}
	helper(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

