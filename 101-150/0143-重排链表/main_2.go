package _143_重排链表

func reorderList2(head *ListNode)  {
	if head == nil || head.Next == nil{
		return
	}
	midNode := splitList2(head)
	midNext := midNode.Next
	midNode.Next = nil
	l2, _ := reverseList2(midNext)
	head = mergeList(head, l2)
	return
}

func splitList2(head *ListNode) (midNode *ListNode){
	pre := &ListNode{Next:head}
	slow, fast := pre, pre
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}


func reverseList2(head *ListNode) (newHead, newTail *ListNode){
	var pre *ListNode
	curr := head
	for curr != nil{
		pre, curr, curr.Next = curr, curr.Next, pre
	}
	return pre, head
}

func mergeList(l0, l1 *ListNode) *ListNode{
	head := &ListNode{}
	res := head
	for {
		if l0 == nil && l1 == nil{
			break
		}
		if l0 == nil{
			res.Next = l1
			break
		}else{
			res.Next = l0
			l0 = l0.Next
			res = res.Next
		}
		if l1 == nil{
			res.Next = l0
			break
		}else{
			res.Next = l1
			l1 = l1.Next
			res = res.Next
		}
	}
	return head.Next
}
