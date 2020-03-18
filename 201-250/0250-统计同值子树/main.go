package _250_统计同值子树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
	1. 叶结点
	2. 只有左结点，且值相同
	3. 只有右结点，且值相同
	4. 有左右结点，满足2，3
*/

func countUnivalSubtrees(root *TreeNode) int {
	count := 0
	isSame(root, &count)
	return count
}

func isSame(root *TreeNode, count *int)  bool{
	if root == nil{
		return true
	}
	leftOk := false
	if root.Left == nil || isSame(root.Left, count) && root.Left.Val == root.Val{
		leftOk = true
	}

	rightOk := false
	if root.Right == nil || isSame(root.Right, count) && root.Right.Val == root.Val{
		rightOk = true
	}

	if leftOk && rightOk{
		*count++
		return true
	}
	return false
}
