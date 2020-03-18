package _148_排序链表

type ListNode struct {
	Val int
	Next *ListNode
}

/*
	归并排序 O(nlogn)
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	left, right := splitList(head)
	return merge(sortList(left), sortList(right))
}

func splitList(head *ListNode) (left, right *ListNode){
	if head == nil || head.Next == nil{
		return
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	left = head
	right = slow.Next
	slow.Next = nil
	return
}

func merge(l0, l1 *ListNode) *ListNode{
	pre := &ListNode{}
	curr := pre
	for {
		if l0 == nil && l1 == nil{
			break
		}
		if l0 == nil{
			curr.Next = l1
			break
		}
		if l1 == nil{
			curr.Next = l0
			break
		}
		if l0.Val < l1.Val{
			curr.Next = l0
			l0 = l0.Next
		}else{
			curr.Next = l1
			l1 = l1.Next
		}
		curr = curr.Next
	}
	return pre.Next
}
