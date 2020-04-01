package _876_链表的中间结点

import (
	"fmt"
	"testing"
)

func TestMiddleNode(t *testing.T){
	fmt.Println(middleNode(mockL1()).Val)
	fmt.Println(middleNode(mockL2()).Val)
}

func mockL1() *ListNode{
	return &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
}

func mockL2() *ListNode{
	return &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
			},
		},
	}
}