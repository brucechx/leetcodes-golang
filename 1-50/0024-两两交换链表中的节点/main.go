package _024_两两交换链表中的节点

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}


// 递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// node to swap
	firstNode := head
	secondNode := head.Next
	// swapping
	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode
	// new head node
	return secondNode
}


// 迭代
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	pre := &ListNode{}
	pre.Next = head
	preNode := pre
	for head != nil && head.Next != nil{
		// to swap
		firstNode := head
		secondNode := head.Next

		// swapping
		preNode.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		// new
		preNode = firstNode
		head = firstNode.Next
	}
	return pre.Next
}