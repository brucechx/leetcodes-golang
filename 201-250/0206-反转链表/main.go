package _206_反转链表

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil{
		pre, curr, curr.Next = curr, curr.Next, pre
	}
	return pre
}
