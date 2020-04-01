package _876_链表的中间结点


type ListNode struct {
	Val int
 	Next *ListNode
}

/*
	快慢指针
*/

func middleNode(head *ListNode) *ListNode {
	//slow, fast := &ListNode{Next:head}, &ListNode{Next:head} // 这个显示前面一个
	slow, fast := head, head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}