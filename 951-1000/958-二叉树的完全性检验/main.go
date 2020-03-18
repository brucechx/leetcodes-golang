package _58_二叉树的完全性检验

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
	按层遍历，如果有中间结点为空，但后续结点不为空即不是完全二叉树
*/

func isCompleteTree(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	flag := false
	for len(queue) > 0 {
		nextQueue := make([]*TreeNode, 0)
		for _, node := range queue{
			if node == nil{
				flag = true
				continue
			}
			if flag {
				return false
			}
			nextQueue = append(nextQueue, node.Left, node.Right)
		}
		queue = nextQueue
	}
	return true
}

