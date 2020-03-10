package _083_删除排序链表中的重复元素

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	current := head
	for current.Next != nil{
		if current.Val == current.Next.Val{
			current.Next = current.Next.Next
		}else {
			current = current.Next
		}
	}
	return head
}
