package _141_环形链表

type ListNode struct {
	Val int
	Next *ListNode
}

// 快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil{
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil && slow != fast{
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow == fast
}
