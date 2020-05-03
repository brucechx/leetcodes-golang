package _021_合并两个有序链表

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T)  {
	cases := []struct{
		l1 *ListNode
		l2 *ListNode
		res string
	}{
		{initl1(), initl2(), "112344"},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, echoList(mergeTwoLists(cas.l1, cas.l2)))
	}
}

func initl1() *ListNode{
	n2 := &ListNode{Val:4}
	n1 := &ListNode{Val:2, Next:n2}
	head := &ListNode{Val:1, Next:n1}
	return head
}

func initl2() *ListNode{
	n2 := &ListNode{Val:4}
	n1 := &ListNode{Val:3, Next:n2}
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
