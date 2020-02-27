package _107_二叉树的层次遍历_II


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil{
		return res
	}
	currentLevelNodes := []*TreeNode{root}
	for len(currentLevelNodes) != 0 {
		var levelNodes []int
		var nextLevelNodes []*TreeNode
		for _, node := range currentLevelNodes{
			levelNodes = append(levelNodes, node.Val)
			if node.Left != nil{
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil{
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}
		res = append(res, levelNodes)
		currentLevelNodes = nextLevelNodes
	}
	for i:=0; i<len(res)/2; i++{
		res[i], res[len(res) - i - 1] = res[len(res) - i - 1], res[i]
	}
	return res
}
