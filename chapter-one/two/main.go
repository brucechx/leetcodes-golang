package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type ListNode struct {
	    Val int
	    Next *ListNode
	}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:0,
	}
	carry := 0
	curr := head

	for{
		if nil == l1 && nil == l2{
			break
		}
		var x, y int
		if nil != l1{
			x = l1.Val
		}
		if nil != l2{
			y = l2.Val
		}
		sum := x + y + carry
		carry = sum / 10
		tmp := sum % 10
		curr.Next = &ListNode{
			Val:tmp,
		}
		curr = curr.Next
		if nil != l1{
			l1 = l1.Next
		}
		if nil != l2{
			l2 = l2.Next
		}
	}
	if carry == 1{
		curr.Next = &ListNode{
			Val:carry,
		}
	}
	return head.Next
}


func createList(size int) *ListNode{
	if size <= 0{
		return &ListNode{}
	}
	head := &ListNode{
		Val:0,
	}
	curr := head

	//r := rand.New(rand.NewSource(time.Now().Unix()))
	for i:=0; i< size; i++{
		//tmp := rand.Intn(10)
		//tmp := r.Intn(10)
		tmp := <-rand_generator(10)
		node := &ListNode{
			Val:tmp,
		}
		curr.Next = node
		curr = node
	}
	return head.Next
}

func rand_generator(n int) chan int {
	rand.Seed(time.Now().UnixNano())
	out := make(chan int)
	go func(x int) {
		for {
			out <- rand.Intn(x)
		}
	}(n)
	return out
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

func main() {
	l1 := createList(3)
	l2 := createList(3)
	l1Str := printList(l1)
	l2Str := printList(l2)

	fmt.Println("l1: ", l1Str)
	fmt.Println(" + ")
	fmt.Println("l2: ", l2Str)

	resultList := addTwoNumbers(l1, l2)
	resultStr := printList(resultList)
	fmt.Println("=")
	fmt.Println(resultStr)

}
