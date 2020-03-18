package _148_排序链表

import (
	"fmt"
	"testing"
)

func TestSortList(t *testing.T){
	head := mockList()
	fmt.Println(echoList(sortList(head)))
}

func mockList() *ListNode{
	return &ListNode{Val:4, Next:&ListNode{Val:2, Next:&ListNode{Val:1, Next:&ListNode{Val:3}}}}
}

func echoList(head *ListNode) string{
	var res string
	for head != nil{
		res += fmt.Sprintf("%d", head.Val)
		head = head.Next
	}
	return res
}

