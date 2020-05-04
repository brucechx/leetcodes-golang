package _637_二叉树的层平均值

type TreeNode struct {
	Val int
 	Left *TreeNode
  	Right *TreeNode
}

// 按层遍历，然后求每层平均值
func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sum, count := 0, 0
		nextQueue := make([]*TreeNode, 0)
		for len(queue) > 0{
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			count++
			if node.Left != nil{
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil{
				nextQueue = append(nextQueue, node.Right)
			}
		}
		coverage := float64(sum) / float64(count)
		res = append(res, coverage)
		queue = nextQueue
	}
	return res
}
