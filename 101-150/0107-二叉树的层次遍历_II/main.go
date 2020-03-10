package _107_二叉树的层次遍历_II


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil{
		return [][]int{}
	}
	currentLevel := make([]*TreeNode, 0)
	currentLevel = append(currentLevel, root)
	res := make([][]int, 0)
	for len(currentLevel) != 0 {
		nextLevel := make([]*TreeNode, 0)
		currentLevelVal := make([]int, 0)
		for _, node := range currentLevel{
			currentLevelVal = append(currentLevelVal, node.Val)
			if node.Left != nil{
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil{
				nextLevel = append(nextLevel, node.Right)
			}
		}
		currentLevel = nextLevel
		res = append([][]int{currentLevelVal}, res...)
	}
	return res
}
