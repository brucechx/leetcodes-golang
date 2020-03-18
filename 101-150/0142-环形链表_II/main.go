package _142_环形链表_II

type ListNode struct {
	Val int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return nil
	}
	slow, fast := head.Next, head.Next.Next
	//slow, fast := head, head.Next  ==> slow.Next == head
	for fast != nil && fast.Next != nil && slow != fast{
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow != fast{
		return nil
	}
	// slow.Next == head
	for slow != head{
		slow, head = slow.Next, head.Next
	}
	return head
}
