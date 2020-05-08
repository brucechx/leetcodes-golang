package _662_二叉树最大宽度

type TreeNode struct {
	Val int
    Left *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	left := make(map[int]int)
	max := 0
	var dfs func(root *TreeNode, depth, pos int)
	dfs = func(root *TreeNode, depth, pos int) {
		if root == nil{
			return
		}
		if _, ok := left[depth]; !ok{
			left[depth] = pos // 其一次遍历的起始坐标
		}
		max = maxN(max, pos - left[depth] + 1)
		dfs(root.Left, depth+1, 2*pos)
		dfs(root.Right, depth+1, 2*pos+1)
	}
	dfs(root, 0, 0)
	return max
}

func maxN(a, b int) int {
	if a > b {
		return a
	}
	return b
}