package _024_两两交换链表中的节点

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestSwapPairs(t *testing.T){
	cases := []struct{
		l1 *ListNode
		res string
	}{
		{initl1(), "2143"},
		{initl2(), "213"},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, echoList(swapPairs2(cas.l1)))
	}
}

func initl1() *ListNode{
	n3 := &ListNode{Val:4}
	n2 := &ListNode{Val:3, Next:n3}
	n1 := &ListNode{Val:2, Next:n2}
	head := &ListNode{Val:1, Next:n1}
	return head
}

func initl2() *ListNode{
	n2 := &ListNode{Val:3}
	n1 := &ListNode{Val:2, Next:n2}
	head := &ListNode{Val:1, Next:n1}
	return head
}

func echoList(head *ListNode) string{
	var res string
	for head != nil{
		res += fmt.Sprintf("%d", head.Val)
		head = head.Next
	}
	return res
}