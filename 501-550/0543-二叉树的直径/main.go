package _543_二叉树的直径

type TreeNode struct {
 	Val int
 	Left *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_ , res := helper(root)
	return res
}

func helper(root *TreeNode) (depth, diameter int){
	if root == nil{
		return
	}
	lDepth, lDiameter := helper(root.Left)
	rDepth, rDiameter := helper(root.Right)
	depth = max(lDepth, rDepth) + 1
	diameter = max(lDepth + rDepth, max(lDiameter, rDiameter))
	return
}

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}
