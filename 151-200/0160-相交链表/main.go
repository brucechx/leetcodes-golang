package _160_相交链表

type ListNode struct {
	Val int
	Next *ListNode
}

/*
 	入栈法 或者 交替两次遍历
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	hasLinkedToB, hasLinkedToA := false, false
	for a != nil && b != nil{
		if a == b{
			return b
		}
		a, b = a.Next, b.Next
		if a == nil && ! hasLinkedToB{
			a = headB
			hasLinkedToB = true
		}
		if b == nil && ! hasLinkedToA{
			b = headA
			hasLinkedToA = true
		}
	}
	return nil
}
