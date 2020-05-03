package _337_打家劫舍_III

// 二叉树中不相邻节点之间的最大和
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
1。对一个节点；有两种状态，偷(1)或者不偷(0)
2。不能相邻；则
	当前节点不偷的时候： 最大结果 = 当前左孩子偷时的金额 + 当前右孩子偷时的金额
	当前节点偷的时候； 最大结果 = 当前左孩子不偷时的最大金额 + 当前右孩子不偷时的最大金额 + 当前节点

 */
func rob(root *TreeNode) int {
	result := robInternal(root)
	return max(result[0], result[1])
}

func robInternal(root *TreeNode) [2]int{
	result := [2]int{}
	if root == nil{
		return result
	}
	left := robInternal(root.Left)
	right := robInternal(root.Right)
	// 不偷
	result[0] = max(left[0], left[1]) + max(right[0], right[1])
	// 偷
	result[1] = left[0] + right[0] + root.Val
	return result
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}
