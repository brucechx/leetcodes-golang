package _222_完全二叉树的节点个数

/*
	时间复杂度：O(N)
	空间复杂度: O(d) = O(log N),其中 dd 指的是树的的高度，运行过程中堆栈所使用的空间。
 */
func countNodes2(root *TreeNode) int {
	return helper2(root)
}

func helper2(root *TreeNode) int{
	if root == nil{
		return 0
	}
	return 1 + helper2(root.Left) + helper2(root.Right)
}
