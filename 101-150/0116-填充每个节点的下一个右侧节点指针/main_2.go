package _116_填充每个节点的下一个右侧节点指针


// bfs
func connect2(root *Node) *Node {
	if root == nil{
		return nil
	}
	currentQueue := make([]*Node, 0)
	currentQueue = append(currentQueue, root)
	for len(currentQueue) > 0{ // 每层
		nextQueue := make([]*Node, 0)
		pre := &Node{}
		for i:=0; i<len(currentQueue); i++{
			curr := currentQueue[0]
			if i > 0{
				pre.Next = curr // 右指向
			}
			pre = curr
			if curr.Left != nil{
				nextQueue = append(nextQueue, curr.Left) // 入队列，下一层
			}
			if curr.Right != nil{
				nextQueue = append(nextQueue, curr.Right)
			}
		}
		currentQueue = nextQueue
	}
	return root
}
