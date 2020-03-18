package _156_上下翻转二叉树

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

/*
	左子结点变为父节点； 右结点变左结点； 父节点变为右子结点
*/

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	var parent, parentRight *TreeNode
	for root != nil{
		rootLeft := root.Left
		root.Left = parentRight
		parentRight = root.Right
		root.Right = parent
		parent = root
		root = rootLeft
	}
	return parent
}
