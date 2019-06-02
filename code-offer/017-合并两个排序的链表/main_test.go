package _17_合并两个排序的链表

import (
	"testing"
	"reflect"
)

func initLinkList() *Node{
	lastNode := Node{
		3,
		nil,
	}
	secondNode := Node{
		2,
		&lastNode,
	}
	firstNode := Node{
		1,
		&secondNode,
	}
	return &firstNode
}

func TestMerge(t *testing.T) {
	l1FourthNode := &Node{20, nil}
	l1ThirdNode := &Node{15, l1FourthNode}
	l1SecondNode := &Node{10, l1ThirdNode}
	l1 := &Node{5, l1SecondNode}

	l2ThirdNode := &Node{20, nil}
	l2SecondNode := &Node{3, l2ThirdNode}
	l2 := &Node{2, l2SecondNode}
	gotNode := SortedMerge(l1, l2)
	var got []int
	for gotNode != nil {
		got = append(got, gotNode.val)
		gotNode = gotNode.next
	}
	want := []int{2, 3, 5, 10, 15, 20, 20}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestMerge2(t *testing.T) {
	l1FourthNode := &Node{20, nil}
	l1ThirdNode := &Node{15, l1FourthNode}
	l1SecondNode := &Node{10, l1ThirdNode}
	l1 := &Node{5, l1SecondNode}

	l2ThirdNode := &Node{20, nil}
	l2SecondNode := &Node{3, l2ThirdNode}
	l2 := &Node{2, l2SecondNode}
	gotNode := SortedMerge2(l1, l2)
	var got []int
	for gotNode != nil {
		got = append(got, gotNode.val)
		gotNode = gotNode.next
	}
	want := []int{2, 3, 5, 10, 15, 20, 20}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}
