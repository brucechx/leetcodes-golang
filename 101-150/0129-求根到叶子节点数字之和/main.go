package _129_求根到叶子节点数字之和

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	res := 0
	var dfs func(root *TreeNode, tmp int)
	dfs = func(root *TreeNode, tmp int) {
		if root == nil{
			return
		}

		tmp = tmp * 10 + root.Val

		if root.Left == nil && root.Right == nil{
			res += tmp
			return
		}

		dfs(root.Left, tmp)
		dfs(root.Right, tmp)
	}
	dfs(root, 0)
	return res
}
