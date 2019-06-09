package _56_链表中环的入口结点

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLinkList_CircleNode(t *testing.T) {
	l := NewLinkList()
	l.append(1)
	l.append(2)
	l.append(3)
	l.append(8)
	l.append(5)
	l.append(7)
	l.append(6)

	node1 := l.getNode(8)
	node2 := l.getNode(6)
	node2.next = node1
	/* 1 --> 2 -->3 -->8 --> 5
	                  |      |
	                  6  <-- 7
	*/
	assert.Equal(t, true, l.circle(l.head))
	assert.Equal(t, node1, l.CircleNode(l.head))
}
