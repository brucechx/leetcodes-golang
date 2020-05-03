package _019_删除链表的倒数第N个节点

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T){
	cases := []struct{
		head *ListNode
		n int
		res string
	}{
		{initList(), 1, "012"},
		{initList(), 2, "013"},
		{initList(), 4, "123"},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, echoList(removeNthFromEnd(cas.head, cas.n)))
	}
}

func initList() *ListNode{
	n3 := &ListNode{Val:3}
	n2 := &ListNode{Val:2, Next:n3}
	n1 := &ListNode{Val:1, Next:n2}
	head := &ListNode{Val:0, Next:n1}
	return head
}

func echoList(head *ListNode) string{
	node := head
	var res string
	for node != nil{
		res += fmt.Sprintf("%d", node.Val)
		node = node.Next
	}
	return res
}

