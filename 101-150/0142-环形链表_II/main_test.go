package _142_环形链表_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectCycle(t *testing.T){
	cases := []struct{
		head *ListNode
		res int
	}{
		{
			head: mockList(true),
			res: 2,
		},
		//{
		//	head: mockList(false),
		//	res: -1,
		//},
	}
	for _, cas := range cases{
		res := detectCycle(cas.head)
		result := -1
		if res != nil{
			result = res.Val
		}
		assert.Equal(t, cas.res, result)
	}

}

func mockList(isCycle bool) *ListNode{
	head := &ListNode{Val:3}
	n1 := &ListNode{Val:2}
	n2 := &ListNode{Val:0}
	n3 := &ListNode{Val:-4}
	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	if isCycle{
		n3.Next = n1
	}
	return head
}
