package _57_删除链表中重复的结点

import "fmt"

//链表结点
type Node struct{
	data interface{}
	next *Node
}

//链表
type LinkList struct{
	head *Node
	tail *Node
	size int
}

func NewList() *LinkList{
	return &LinkList{}
}

//在链表尾部添加数据
func(l *LinkList) append(e interface{}){
	//为数据创建新结点
	newNode:=Node{}
	newNode.data=e
	newNode.next=nil

	if l.size==0 {
		l.head=&newNode
		l.tail=&newNode
	}else{
		l.tail.next=&newNode
		l.tail=&newNode
	}
	l.size++
}

func (l *LinkList) deleteDuplicate() *Node{
	if l.head == nil{
		return nil
	}

	// 临时的头节点
	root := &Node{}
	root.next = l.head
	// 记录前驱结点
	pre := root
	// 记录当前处理的结点
	curr := l.head
	for curr != nil && curr.next != nil{
		// 有重复结点，与curr值相同的结点都要删除
		if curr.data == curr.next.data{
			// 找到下一个不同值的节点，注意其有可能也是重复节点
			for curr.next!=nil && curr.next.data == curr.data{
				curr = curr.next
			}
			pre.next = curr.next
		}else {
			pre.next = curr
			pre = pre.next
		}
		curr = curr.next
	}
	return root.next
}

func printList(node *Node) string{
	var res string
	for node != nil{
		res += fmt.Sprintf("%v", node.data)
		node = node.next
	}
	return res
}