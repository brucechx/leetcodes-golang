package _445_两数相加_II

type ListNode struct {
	Val int
	Next *ListNode
}

/*
	不能反转链表，则双栈法
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := make([]int, 0), make([]int, 0)
	for l1 != nil{
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil{
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	sum := 0
	head := &ListNode{}
	for len(stack1) >0 || len(stack2)>0{
		// stack1 pop
		if len(stack1) > 0 {
			sum += stack1[len(stack1) - 1]
			stack1 = stack1[:len(stack1) - 1]
		}
		// stack2 pop
		if len(stack2) > 0 {
			sum += stack2[len(stack2) - 1]
			stack2 = stack2[:len(stack2) - 1]
		}

		head.Val = sum % 10 // 当前点
		tmp := &ListNode{ // 向前进位
			Val:  sum / 10 ,
			Next: head,
		}
		head = tmp

		sum /= 10
	}
	if head.Val == 0{
		return head.Next
	}
	return head
}
