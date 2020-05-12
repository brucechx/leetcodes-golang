package _783_二叉搜索树节点最小距离

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 同530题，中序遍历
func minDiffInBST(root *TreeNode) int {
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
