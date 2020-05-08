package _669_修剪二叉搜索树

type TreeNode struct {
	Val int
    Left *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil{
		return root
	}
	if root.Val > R { // 太大
		return trimBST(root.Left, L, R)
	}
	if root.Val < L { // 太小
		return trimBST(root.Right, L, R)
	}
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}
