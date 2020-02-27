package _061_旋转链表

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T){
	cases := []struct{
		l1 *ListNode
		k int
		res string
	}{
		{initl1(), 1, "1"},
		{initl2(), 4, "201"},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, echoList(rotateRight(cas.l1, cas.k)))
	}
}

func initl1() *ListNode{
	head := &ListNode{Val:1}
	return head
}

func initl2() *ListNode{
	n2 := &ListNode{Val:2}
	n1 := &ListNode{Val:1, Next:n2}
	head := &ListNode{Val:0, Next:n1}
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
