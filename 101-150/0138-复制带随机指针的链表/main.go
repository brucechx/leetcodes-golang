package _138_复制带随机指针的链表

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		temp := &Node{cur.Val, nil, nil}
		m[cur] = temp
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]

}
