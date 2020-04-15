package _116_填充每个节点的下一个右侧节点指针

type Node struct {
      Val int
      Left *Node
      Right *Node
      Next *Node
}

func connect(root *Node) *Node {
	if root == nil{
		return nil
	}
	start := root // 每层的开始结点
	curr := &Node{} // 游标结点
	for start.Left != nil{ // 每层遍历
		curr = start
		for curr != nil  { // 当前层
			curr.Left.Next = curr.Right // 兄弟结点
			if curr.Next != nil{ // 兄弟树
				curr.Right.Next = curr.Next.Left
			}
			curr = curr.Next
		}
		start = start.Left // 每层遍历
	}
	return root
}
