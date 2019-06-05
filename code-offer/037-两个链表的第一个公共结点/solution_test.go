package _37_两个链表的第一个公共结点

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindFirstCommonNode(t *testing.T) {
	head1, head2 := initList()
	cases := []struct{
		head1 *Node
		head2 *Node
		res int
	}{
		{head1, head2, 11},
	}
	for _, cas := range cases{
		assert.Equal(t, cas.res, FindFirstCommonNode(cas.head1, cas.head2))
	}
}

func initList() (head1, head2 *Node){
	common := &Node{
		11, &Node{
			21, &Node{
				31, &Node{
					41, &Node{
						51, nil}}}}}
	head1 = &Node{
		1, &Node{
			2, &Node{
				3, &Node{
					4, &Node{
						5, common}}}}}
	head2 = &Node{
		10, &Node{
			20, &Node{
				30, &Node{
					40, common}}}}
	return
}
