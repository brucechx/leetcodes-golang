package _39_二叉树的深度

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTreeDepth(t *testing.T) {
	n := initTree()
	res := TreeDepth(n)
	assert.Equal(t, res, 4)
}

/*
			4
		2		  7
	  1   3		6	8
              5
*/
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
