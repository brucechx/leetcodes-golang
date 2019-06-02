package _23_从上往下打印二叉树

import "testing"

func TestPrintFromTopToBottom(t *testing.T) {
	/*
				1
	          /   \
	         2     3
	        / \   /
	       4   5 6
	      /     /
	     7     8
	 */
	n1 := &Node{Data: 1}
	n2 := &Node{Data: 2}
	n3 := &Node{Data: 3}
	n4 := &Node{Data: 4}
	n5 := &Node{Data: 5}
	n6 := &Node{Data: 6}
	n7 := &Node{Data: 7}
	n8 := &Node{Data: 8}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n4.Left = n7
	n6.Left = n8

	PrintFromTopToBottom(n1)

//	output:
//   	1 2 3 4 5 6 7 8
}