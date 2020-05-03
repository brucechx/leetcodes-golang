package __遍历二叉树的神级方法

import "fmt"

type Node struct {
	Left *Node
	Right *Node
	Val int
}

func inorder(root *Node){
	if root == nil{
		return
	}
	curr1 := root
	var curr2 *Node
	for curr1 != nil{
		curr2 = curr1.Left
		if curr2 != nil{
			// 左子树的最右节点
			for curr2.Right != nil && curr2.Right != curr1{
				curr2 = curr2.Right
			}
			// 最右节点，指向当前节点
			if curr2.Right == nil{
				curr2.Right = curr1
				curr1 = curr1.Left
				continue
			}else {
				// 复位
				curr2.Right = nil
			}
		}
		fmt.Println("res=", curr1.Val)
		curr1 = curr1.Right
	}
}
