package _023_合并K个排序链表

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMergeKList(t *testing.T){
	cases := []struct{
		l []*ListNode
		res string
	}{
		{[]*ListNode{initl1(), initl2()}, "112344"},
	}

	for _, cas := range cases{
		assert.Equal(t, cas.res, echoList(merge(cas.l)))
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