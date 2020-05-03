package _019_删除链表的倒数第N个节点

import "fmt"

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	parentNode, isHead := getParentNode(head, n)
	if isHead{
		return head.Next
	}
	fmt.Println("parentNode=", parentNode.Val, " isHead=", isHead)
	parentNode.Next = parentNode.Next.Next
	return head
}

func getParentNode(head *ListNode, n int) (parentNode *ListNode, isHead bool){
	curr := head
	parentNode = head
	for curr != nil{
		if n >= 0 {
			curr = curr.Next
			n--
			continue
		}
		parentNode = parentNode.Next
		curr = curr.Next
	}
	return parentNode, n == 0
}