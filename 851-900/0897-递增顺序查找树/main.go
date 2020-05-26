package _897_递增顺序查找树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 1. 中序遍历 + 构造新的树
// 2. 中序遍历 + 更改树的连接方式
func increasingBST(root *TreeNode) *TreeNode {
	pre := &TreeNode{}
	curr := pre
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil{
			return
		}
		inOrder(node.Left)
		node.Left = nil
		curr.Right = node
		curr = node
		inOrder(node.Right)
	}
	inOrder(root)
	return pre.Right
}
