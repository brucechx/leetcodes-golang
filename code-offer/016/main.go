package _16

type Node struct {
	val int
	next *Node
}

func ReverseList(head *Node) *Node{
	if head == nil || head.next == nil{
		return nil
	}

	var pre, curr, tmp *Node // tmp 中介节点 作为临时节点
	pre = head  // 前一个节点
	curr = pre.next // 从第二个节点,当前节点开始反转
	for curr != nil{
		tmp = curr.next  // 提取出当前节点的下一个节点，保存到临时节点中
		curr.next = pre // 当前节点的前一个节点作为当前节点的下一个节点，刚好实现反转
		pre = curr // 此时当前节点则为下一次反转的当前节点的前一个节点
		curr = tmp // 当前节点的下一个节点，此时为下一次反转的当前节点
	}
	head.next = nil
	head = pre // 当前节点的前一个节点作为头节点，即最开始的节点
	return head
}

func Reverse(head *Node) *Node{
	if head == nil || head.next == nil{
		return head
	}
	cur := head
	var pre *Node
	for cur != nil{
		pre, cur, cur.next = cur, cur.next, pre
	}
	head = pre
	return head
}
