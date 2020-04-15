package _117_填充每个节点的下一个右侧节点指针_II

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}


func connect(root *Node) *Node {
	head := root
	for head != nil{ // 当前层头结点
		// 处理当前层
		curr := head // 当前结点
		var nextHead, nextTail *Node // 下一层头结点, // 下一层游标结点
		for curr != nil{ // 当前层处理
			if curr.Left != nil{
				if nextTail == nil{
					nextTail = curr.Left
					nextHead = curr.Left
				}else {
					nextTail.Next = curr.Left
					nextTail = nextTail.Next
				}
			}
			if curr.Right != nil{
				if nextTail == nil{
					nextTail = curr.Right
					nextHead = curr.Right
				}else {
					nextTail.Next = curr.Right
					nextTail = nextTail.Next
				}
			}
			curr = curr.Next
		}
		head = nextHead
	}
	return root
}
