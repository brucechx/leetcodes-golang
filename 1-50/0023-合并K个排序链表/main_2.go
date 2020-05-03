package _023_合并K个排序链表

import "fmt"

// 两两合并
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0{
		return nil
	}
	return mergeLists(lists)
}

func mergeLists(lists []*ListNode) *ListNode{
	listLen := len(lists)
	half := listLen / 2
	if listLen == 0 {
		return nil
	}
	if listLen == 1{
		return lists[0]
	}
	if listLen != 2{
		return mergeKLists([]*ListNode{mergeKLists(lists[:half]), mergeKLists(lists[half:])})
	}
	return mergeTwoList(lists[0], lists[1])

}

func mergeTwoList(l0, l1 *ListNode) *ListNode{
	head := &ListNode{}
	res := head
	for {
		if l0 == nil && l1 == nil{
			break
		}
		if l0 == nil{
			res.Next = l1
			break
		}
		if l1 == nil{
			res.Next = l0
			break
		}
		if l0.Val < l1.Val {
			res.Next = l0
			l0 = l0.Next
		}else{
			res.Next = l1
			l1 = l1.Next
		}
		res = res.Next
	}
	return head.Next
}

func echoList2(head *ListNode) string{
	var res string
	for head != nil{
		res += fmt.Sprintf("%d", head.Val)
		head = head.Next
	}
	return res
}