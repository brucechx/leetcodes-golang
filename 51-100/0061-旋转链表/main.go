package _061_旋转链表

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	// 链表环
	oldTail := head
	listLen := 1
	for oldTail.Next != nil{
		oldTail = oldTail.Next
		listLen++
	}
	// 构成环
	oldTail.Next = head

	// 寻新的尾结点：n - k % n - 1
	// 新的头结点：n - k % n
	newTail := head
	for i:=0; i< listLen - k%listLen - 1; i++{
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}

func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil{
		return head
	}
	nodeLen, TailNode := getNodeLen(head)
	if nodeLen <= k{
		k = k % nodeLen
	}
	if k == 0{
		return head
	}
	kkk := nodeLen - k
	KPNode := getKParentNode(head, kkk)
	fmt.Println("KPNode=",KPNode.Val, " nodeLen=",nodeLen, " k=",k, " kkk=", k)
	tmp := KPNode.Next
	KPNode.Next = nil
	TailNode.Next = head
	head = tmp
	//head, TailNode.Next, KPNode.Next = KPNode.Next, head, nil
	return head
}

func getNodeLen(head *ListNode) (length int, TailNode *ListNode){
	i := 1
	for head.Next != nil{
		i++
		head = head.Next
	}
	return i, head
}

func getKParentNode(head *ListNode, k int) (KParentNode *ListNode){
	for k > 1 && head != nil{
		head = head.Next
		k--
	}
	return head
}
