package _025_K个一组翻转链表

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)


func TestSwapPairs(t *testing.T){
	cases := []struct{
		l1 *ListNode
		k int
		res string
	}{
		{initl1(), 3, "321654"},
		{initl1(), 4, "432156"},
		{initl2(), 2, "213"},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, echoList(reverseKGroup(cas.l1, cas.k)))
	}
}

func initl1() *ListNode{
	n5 := &ListNode{Val:6}
	n4 := &ListNode{Val:5, Next:n5}
	n3 := &ListNode{Val:4, Next:n4}
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
