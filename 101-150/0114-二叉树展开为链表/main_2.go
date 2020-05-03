package _114_二叉树展开为链表


func flatten2(root *TreeNode)  {
	recur(root)
}

func recur(root *TreeNode) *TreeNode{
	if root == nil ||
		(root.Left == nil && root.Right == nil) {
		return root
	}

	if root.Left == nil {
		return recur(root.Right)
	}

	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return recur(root.Right)
	}

	res := recur(root.Right)
	recur(root.Left).Right = root.Right
	root.Right = root.Left
	// 不要忘记封闭 left 节点
	root.Left = nil

	return res
}