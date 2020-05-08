package _701_二叉搜索树中的插入操作

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{Val:val}
	if root == nil{
		return newNode
	}
	insert(root, newNode)
	return root
}

func insert(curr, new *TreeNode) {
	if curr.Val > new.Val{
		if curr.Left == nil{
			curr.Left = new
		}else {
			insert(curr.Left, new)
		}
	}else {
		if curr.Right == nil{
			curr.Right = new
		}else {
			insert(curr.Right, new)
		}
	}
}
