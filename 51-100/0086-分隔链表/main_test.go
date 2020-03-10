package _086_分隔链表

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T){
	mocklist := mockList()
	newHead := partition2(mocklist, 3)
	for newHead != nil{
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}

}


func mockList() *ListNode{
	head := &ListNode{Val:1}
	head.Next = &ListNode{Val:4}
	head.Next.Next = &ListNode{Val:3}
	head.Next.Next.Next = &ListNode{Val:2,
		Next:&ListNode{
			Val:  5,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		}}
	return head

}