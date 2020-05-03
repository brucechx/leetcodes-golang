package _021_合并两个有序链表


func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	pre := &ListNode{}
	curr := pre
	for {
		if l1 == nil && l2 == nil{
			break
		}

		if l1 == nil{
			curr.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
			curr = curr.Next
			continue
		}
		if l2 == nil{
			curr.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
			curr = curr.Next
			continue
		}

		if l1.Val < l2.Val{
			curr.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
			curr = curr.Next
		}else{
			curr.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
			curr = curr.Next
		}
	}
	return pre.Next
}
