package _143_重排链表

import "testing"

func TestReorderList(t *testing.T){
	cases := []struct{
		head *ListNode
	}{
		{
			head: mockList(),
		},
	}
	for _, cas := range cases{
		reorderList(cas.head)
	}
}

func mockList() *ListNode{
	return &ListNode{Val:1, Next:&ListNode{Val:2, Next:&ListNode{Val:3, Next:&ListNode{Val:4}}}}
}
