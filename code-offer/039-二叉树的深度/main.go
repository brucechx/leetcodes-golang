package _39_二叉树的深度

type Node struct {
	Data int
	Left *Node
	Right *Node
}

func TreeDepth(node *Node) (int){
	if node == nil{
		return 0
	}
	result := 1
	l := TreeDepth(node.Left)
	r := TreeDepth(node.Right)
	if l > r {
		result += l
	}else {
		result += r
	}
	return result
}
