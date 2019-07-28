package _002_两数相加

import (
	"fmt"
	"strconv"
)

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	// 定义头节点
	head := &ListNode{}
	curr := head
	// 定义溢出值
	carry := 0
	for {
		if l1 == nil && l2 == nil{
			break
		}
		var x, y int
		if l1 != nil{
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			y = l2.Val
			l2 = l2.Next
		}
		// 从低向高位相加
		sum := x + y + carry
		carry = sum / 10
		val := sum % 10
		curr.Next = &ListNode{
			Val:val,
		}
		curr = curr.Next
	}

	// 高位溢出场景
	if carry == 1{
		curr.Next = &ListNode{
			Val:1,
		}
	}
	return head.Next
}

type ListNode struct {
	Val int
	Next *ListNode
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
