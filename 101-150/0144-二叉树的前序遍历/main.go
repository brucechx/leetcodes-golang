package main

/*
	前序遍历的顺序为根-左-右，具体算法为：
		把根节点push到栈中
		循环检测栈是否为空，若不空，则取出栈顶元素，保存其值
		看其右子节点是否存在，若存在则push到栈中
		看其左子节点，若存在，则push到栈中。
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		if root != nil{
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}else {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack) - 1]
			root = tmp.Right
		}
	}
	return res
}

