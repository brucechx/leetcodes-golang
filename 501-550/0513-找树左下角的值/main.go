package _513_找树左下角的值

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
// 先右后左，层次遍历
*/

func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]
		// 必须先右后左
		if root.Right != nil{
			queue = append(queue, root.Right)
		}
		if root.Left != nil{
			queue = append(queue, root.Left)
		}
	}
	return root.Val
}

