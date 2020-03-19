package _203_移除链表元素

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	preHead := &ListNode{Next: head}
	curr := preHead
	for curr.Next != nil{
		if curr.Next.Val == val{
			curr.Next = curr.Next.Next
		}else{
			curr = curr.Next
		}
	}
	return preHead.Next
}
