package _257_二叉树的所有路径

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil{
		return []string{}
	}
	res := make([]string, 0)
	dfs(root, &res, "")
	return res
}

func dfs(root *TreeNode, res *[]string, path string){
	if root == nil{
		return
	}
	path += fmt.Sprintf("%d", root.Val)
	if root.Left == nil && root.Right == nil{
		*res = append(*res, path)
		return
	}
	path += "->"
	dfs(root.Left, res, path)
	dfs(root.Right, res, path)
}
