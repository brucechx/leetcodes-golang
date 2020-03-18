package _143_重排链表

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

/*
	快慢指针进行二分，后半部分反转，然后合并两链表
*/

func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil{
		return
	}
	halfNode := splitList(head)
	halfHead := halfNode.Next
	halfNode.Next = nil
	halfList, _ := reverseList(halfHead)
	head = merge(head, halfList)
}

func merge(l0, l1 *ListNode) *ListNode{
	res := &ListNode{}
	for {
		if l0 == nil && l1 == nil{
			break
		}
		if l0 == nil{
			res.Next = l1
			break
		}else {
			res.Next = l0
			l0 = l0.Next
			res = res.Next
		}
		if l1 == nil{
			res.Next = l0
			break
		}else {
			res.Next = l1
			l1 = l1.Next
			res = res.Next
		}
	}
	return res.Next
}

func splitList(head *ListNode) (halfNode *ListNode){
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil{
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) (newHead, newTail *ListNode){
	var pre *ListNode
	curr := head
	for curr != nil{
		pre, curr, curr.Next = curr, curr.Next, pre
	}
	return pre, head
}

func echoList(head *ListNode) string{
	var res string
	for head != nil{
		res += fmt.Sprintf("%d", head.Val)
		head = head.Next
	}
	return res
}