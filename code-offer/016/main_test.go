package _16

import (
	"testing"
	"reflect"
)

func TestReverseList(t *testing.T) {
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

	//res := ReverseList(&firstNode)
	res := Reverse(&firstNode)
	var got []int
	for res != nil{
		got = append(got, res.val)
		res = res.next
	}
	want := []int{3, 2, 1}
	if ! reflect.DeepEqual(got, want){
		t.Errorf("want:%v, but get:%v", want, got)
	}
}
