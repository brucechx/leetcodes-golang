package _700_二叉搜索树中的搜索

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil{
		return nil
	}
	if root.Val < val{
		return searchBST(root.Right, val)
	}
	if root.Val > val{
		return searchBST(root.Left, val)
	}
	return root
}

