package _002_两数相加

import (
	"fmt"
	"strconv"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		x *ListNode
		y *ListNode
		expect string
	}{
		{createList([]int{2, 4, 3}), createList([]int{5, 6, 4}), "7 --> 0 --> 8"},
		{createList([]int{2, 4, 3, 1}), createList([]int{5, 6, 4}), "7 --> 0 --> 8 --> 1"},
	}

	for _, cas := range cases{
		res := AddTwoNumbers(cas.x, cas.y)
		assert.Equal(t, cas.expect, res.String() )
	}
}



func createList(l []int) *ListNode{
	if len(l) <= 0{
		return &ListNode{}
	}
	head := &ListNode{
		Val:0,
	}
	curr := head
	for _, v := range l{
		node := &ListNode{
			Val:v,
		}
		curr.Next = node
		curr = node
	}
	return head.Next
}

func printList(node *ListNode) string{
	if nil == node{
		fmt.Println("None")
		return "None"
	}

	res := strconv.Itoa(node.Val)
	node = node.Next
	for{
		if nil == node{
			break
		}
		res += " --> " + strconv.Itoa(node.Val)
		node = node.Next
	}
	return res
}

func (l *ListNode)String() string{
	if nil == l{
		fmt.Println("None")
		return "None"
	}
	node := l
	res := strconv.Itoa(node.Val)
	node = node.Next
	for{
		if nil == node{
			break
		}
		res += " --> " + strconv.Itoa(node.Val)
		node = node.Next
	}
	return res
}