package _147_对链表进行插入排序

type ListNode struct {
	Val int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	newHead := &ListNode{}
	pre := newHead // 插入pre 和 pre.Next 之间
	curr := head  // 待插入的结点
	for curr != nil{
		next := curr.Next
		for pre.Next != nil && pre.Next.Val < curr.Val{
			pre = pre.Next
		}
		// pre.Next.val > curr.Val
		curr.Next = pre.Next
		pre.Next = curr
		// 循环
		pre = newHead
		curr = next
	}
	return newHead.Next
}
