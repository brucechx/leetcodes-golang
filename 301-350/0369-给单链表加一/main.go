package _369_给单链表加一


type ListNode struct {
	Val int
	Next *ListNode
}

/*
	快慢指针；
	找到最靠右边的非9结点；
	- 将其加一
	- 非9结点往后全部置零
	- 所有结点是否都是9?
*/
func plusOne(head *ListNode) *ListNode {
	fast := head
	pre := &ListNode{Next:head, Val:0}
	slow := pre
	for fast != nil{
		if fast.Val != 9{
			slow = fast
		}
		fast = fast.Next
	}
	// 得到最靠右的非9结点
	slow.Val++
	curr := slow.Next
	for curr != nil{
		curr.Val = 0
		curr = curr.Next
	}
	// 所有结点都为9的场景；如999
	if pre.Val == 1{
		return pre
	}
	return head
}
