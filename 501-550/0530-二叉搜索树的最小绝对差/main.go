package _530_二叉搜索树的最小绝对差

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 中序遍历; 二叉搜索树递增
func getMinimumDifference(root *TreeNode) int {
	st := make([]*TreeNode, 0)
	p := root
	pre := -1 << 31
	minVal := 1 << 31 - 1
	for p != nil || len(st) > 0{
		for p != nil{
			st = append(st, p)
			p = p.Left
		}
		// 出栈
		p = st[len(st)-1]
		st = st[:len(st)-1]
		// 处理p
		curr := p.Val
		if curr - pre < minVal{
			minVal = curr - pre
		}
		pre = curr
		// 右节点
		p = p.Right
	}
	return minVal
}
