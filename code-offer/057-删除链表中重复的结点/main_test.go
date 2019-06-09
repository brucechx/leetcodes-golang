package _57_删除链表中重复的结点

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicate(t *testing.T){
	cases := []struct{
		input *LinkList
		res string
	}{
		{initList1(), "125"},
		{initList2(), "3"},
		{initList3(), "12"},
	}

	for _, cas := range cases{
		res := printList(cas.input.deleteDuplicate())
		assert.Equal(t, cas.res, res)
	}

}

func initList1() *LinkList{
	l := NewList()
	l.append(1)
	l.append(2)
	l.append(3)
	l.append(3)
	l.append(4)
	l.append(4)
	l.append(5)
	return l
}

func initList2() *LinkList{
	l := NewList()
	l.append(1)
	l.append(1)
	l.append(2)
	l.append(2)
	l.append(3)
	return l
}

func initList3() *LinkList{
	l := NewList()
	l.append(1)
	l.append(2)
	l.append(3)
	l.append(3)
	l.append(4)
	l.append(4)
	l.append(4)
	return l
}