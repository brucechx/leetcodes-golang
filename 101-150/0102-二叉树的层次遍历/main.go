package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	dfs(root, &res, 0)
	return res
}

func dfs(root *TreeNode, res *[][]int, level int){
	if root == nil {
		return
	}
	if len(*res) == level{
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], root.Val)
	for _, node := range []*TreeNode{root.Left, root.Right}{
		dfs(node, res, level+1)
	}
}
