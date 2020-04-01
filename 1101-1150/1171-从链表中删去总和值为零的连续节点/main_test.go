package _171_从链表中删去总和值为零的连续节点

import (
	"fmt"
	"testing"
)

func TestRemoveZeroSumSublists(t *testing.T){
	l := mockList()
	nl := removeZeroSumSublists(l)
	res := make([]int, 0)
	for c:=nl; c!=nil; c=c.Next{
		res = append(res, c.Val)
	}
	fmt.Println(res)
}

func mockList() *ListNode{
	return &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  -3,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}
}
