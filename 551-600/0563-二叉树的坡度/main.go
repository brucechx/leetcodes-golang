package _563_二叉树的坡度

type TreeNode struct {
	Val int
	Left *TreeNode
    Right *TreeNode
}

func findTilt(root *TreeNode) int {
	tilt := 0
	helper(root, &tilt)
	return tilt
}

func helper(root *TreeNode, tilt *int) (sum int){
	if root == nil{
		return
	}
	lSum := helper(root.Left, tilt)
	rSum := helper(root.Right, tilt)
	sum = lSum + rSum + root.Val
	*tilt += abs(lSum, rSum)
	return
}

func abs(a, b int) int{
	if a - b > 0 {
		return a - b
	}
	return b - a
}

