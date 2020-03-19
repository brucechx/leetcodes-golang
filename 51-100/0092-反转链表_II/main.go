package _092_反转链表_II

type ListNode struct {
   Val int
   Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n{
		return head
	}
	headPre := &ListNode{Next:head}
	m++
	n++
	i := 1
	var mParentNode, mNode, nNode, nNext *ListNode
	currNode := headPre
	for currNode != nil{
		if i == m-1{
			mParentNode = currNode
			mNode = currNode.Next
		}
		if i == n{
			nNext = currNode.Next
			nNode = currNode
		}
		currNode = currNode.Next
		i++
	}
	if mParentNode != nil{
		mParentNode.Next = nil
	}
	if nNode != nil{
		nNode.Next = nil
	}
	newHead, newTail := reverse(mNode)
	if mParentNode != nil{
		mParentNode.Next = newHead
	}
	newTail.Next = nNext
	return headPre.Next
}

func reverse(head *ListNode) (newHead, tail *ListNode){
	var pre *ListNode
	curr := head
	for curr != nil{
		pre, curr, curr.Next = curr, curr.Next, pre
	}
	return pre, head
}
