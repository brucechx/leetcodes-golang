package _236_二叉树的最近公共祖先

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q{
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l == nil{
		// 此时 p 和 q 在 root.Right 中
		return r
	}
	if r == nil{
		// 此时 p 和 q 在 root.Left 中
		return l
	}
	// 此时 p 和 q 分别在 root.Left 和 root.Right 中
	return root
}
