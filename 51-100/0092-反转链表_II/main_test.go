package _092_反转链表_II

import (
	"fmt"
	"testing"
)

func TestReverseBetween(t *testing.T){
	curr := reverseBetween(mockList(), 1, 2)
	fmt.Println(curr.Val, curr.Next.Val)
}

func mockList() *ListNode{
	return &ListNode{
		Val:  3,
		Next: &ListNode{Val:5},
	}
}
