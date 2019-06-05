package _39_平衡二叉树

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsBalanced_Solution(t *testing.T) {
	n := initTree()
	assert.Equal(t, true, IsBalanced_Solution(n))
}

func initTree() *Node{
	n5:= &Node{Data:5}
	n6:= &Node{Data:6, Left:n5}
	n8:= &Node{Data:8}
	n7:= &Node{Data:7, Left:n6, Right:n8}
	n1 := &Node{Data:1}
	n3 := &Node{Data:3}
	n2 := &Node{Data:2, Left:n1, Right:n3}
	n4 := &Node{Data:4, Left:n2, Right:n7}
	return n4
}
