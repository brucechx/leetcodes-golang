package _082_删除排序链表中的重复元素_II

type ListNode struct {
     Val int
     Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	// 如果head 重复, 则删除head
	if head.Val == head.Next.Val{
		for head.Next != nil && head.Val == head.Next.Val{
			head = head.Next
		}
		return deleteDuplicates(head.Next)
	}
	// head 不重复
	head.Next = deleteDuplicates(head.Next)
	return head
}
