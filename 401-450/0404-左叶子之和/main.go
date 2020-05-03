package _404_左叶子之和

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 什么时候遍历到了左叶子节点
// 如何返回记录的左叶子节点的值
var res int

func sumOfLeftLeaves(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode){
	if root == nil{
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil{
		res += root.Left.Val
	}
	dfs(root.Left)
	dfs(root.Right)
	return
}

