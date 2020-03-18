package _222_完全二叉树的节点个数

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	var count int
	helper(root, &count)
	return count
}

func helper(root *TreeNode, count *int){
	if root == nil{
		return
	}
	*count++
	helper(root.Left, count)
	helper(root.Right, count)
}


