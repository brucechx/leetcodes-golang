package _39_平衡二叉树

type Node struct {
	Data int
	Left *Node
	Right *Node
}

func IsBalanced_Solution(pRoot *Node) (result bool) {
	if pRoot == nil {
		return true
	}
	//
	left := TreeDepth3(pRoot.Left)
	right := TreeDepth3(pRoot.Right)
	diff := left - right
	return (diff >= -1 && diff <= 1) && IsBalanced_Solution(pRoot.Left) && IsBalanced_Solution(pRoot.Right)
}

// 使用上一题代码
func TreeDepth3(pRoot *Node) (result int) {
	if pRoot == nil {
		return
	}
	stack := []*Node{pRoot}
	for len(stack) > 0 {
		result++
		tmp := []*Node{}
		for _, v := range stack {
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}

		stack = tmp
	}
	return
}
