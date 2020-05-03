package __遍历二叉树的神级方法

import "testing"

func TestInorder(t *testing.T){
	inorder(mockTree())
}

func mockTree()  *Node{
	root := &Node{
		Left:  &Node{
			Left:  nil,
			Right: nil,
			Val:   2,
		},
		Right: &Node{
			Left:  nil,
			Right: nil,
			Val:   3,
		},
		Val:   1,
	}
	return root
}
