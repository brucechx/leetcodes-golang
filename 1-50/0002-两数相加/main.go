package _002_两数相加

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	// 定义头节点
	head := &ListNode{}
	curr := head
	// 定义溢出值
	carry := 0
	for {
		if l1 == nil && l2 == nil{
			break
		}
		var x, y int
		if l1 != nil{
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			y = l2.Val
			l2 = l2.Next
		}
		// 从低向高位相加
		sum := x + y + carry
		carry = sum / 10
		val := sum % 10
		curr.Next = &ListNode{
			Val:val,
		}
		curr = curr.Next
	}

	// 高位溢出场景
	if carry == 1{
		curr.Next = &ListNode{
			Val:1,
		}
	}
	return head.Next
}
