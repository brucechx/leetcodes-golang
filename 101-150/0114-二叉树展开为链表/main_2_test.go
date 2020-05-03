package _114_二叉树展开为链表

import (
	"fmt"
	"testing"
)

func TestFlatten(t *testing.T){
	root := mockTree()
	flatten(root)
	fmt.Println(echoList(root))
}

func echoList(root *TreeNode) []int{
	res := make([]int, 0)
	for root != nil{
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

func mockTree() *TreeNode{
	root := &TreeNode{
		Val:   0,
		Left:  &TreeNode{Val:1},
		Right: &TreeNode{Val:2},
	}
	return root
}