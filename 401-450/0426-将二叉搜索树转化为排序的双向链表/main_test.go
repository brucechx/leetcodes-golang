package _426_将二叉搜索树转化为排序的双向链表

import (
	"fmt"
	"testing"
)

func TestTreeToDoublyList(t *testing.T){
	doublyList := treeToDoublyList(mockTree())
	fmt.Println(printDoubleList(doublyList)) // 2 3 4 5
	fmt.Println(doublyList.Val) // 2
	fmt.Println(doublyList.Left.Val) // 5
}

func printDoubleList(root *TreeNode) []int{
	p := root
	res := make([]int, 0)
	for p != nil{
		res = append(res, p.Val)
		p = p.Right
		if p == root{
			break
		}
	}
	return res
}

func mockTree() *TreeNode{
	return &TreeNode{
		Val:   4,
		Left:  &TreeNode{
			Val:   3,
			Left:  &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}
}
