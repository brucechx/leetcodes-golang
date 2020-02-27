package _082_删除排序链表中的重复元素_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T){
	cases := []struct{
		head *ListNode
		res []int
	}{
		{
			head: mockList(),
			res: []int{1, 2, 5},
		},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, echo(deleteDuplicates(cas.head)))
	}
}

func mockList() *ListNode{
	l := NewLinkList()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(3)
	l.Append(4)
	l.Append(4)
	l.Append(5)
	return l.head
}

type LinkList struct {
	head *ListNode // 头结点
	tail *ListNode // 尾结点
	size int // 大小
}

//新建空链表，即创建Node指针head，用来指向链表第一个结点，初始为空
func NewLinkList() *LinkList{
	return &LinkList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

//是否为空链表
func (l *LinkList) IsEmpty() bool{
	return l.size == 0
}

//链表长度
func (l *LinkList) Length() int{
	return l.size
}

//在链表尾部添加数据
func(l *LinkList) Append(data int){
	n := &ListNode{
		Val: data,
		Next: nil,
	}
	if l.IsEmpty(){
		l.head = n
		l.tail = n
	}else {
		l.tail.Next = n
		l.tail = n
	}
	l.size++
}

func echo(head *ListNode) []int{
	var res []int
	for head != nil{
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}