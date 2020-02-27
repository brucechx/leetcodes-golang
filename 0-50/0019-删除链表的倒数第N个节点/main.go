package _019_删除链表的倒数第N个节点

// ListNode 是节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode{
	d, headIsNthFromEnd := getDaddy(head, n)

	if headIsNthFromEnd { // 链的长度等于n
		// 删除head节点
		return head.Next
	}

	d.Next = d.Next.Next
	return head
}

// 获取倒数第N个节点的父节点
func getDaddy(head *ListNode, n int) (daddy *ListNode, headIsNthFromEnd bool){
	daddy = head

	for head != nil{
		if n < 0{
			daddy = daddy.Next
		}
		// 先走n步，再一起走，可得倒数第N个点
		n--
		head = head.Next
	}
	// n==0，说明链的长度等于n
	headIsNthFromEnd = n == 0
	return
}
