package _426_将二叉搜索树转化为排序的双向链表

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
	方法一: 中序遍历，最后首尾相链
	方法二: 递归
*/

func treeToDoublyList(root *TreeNode) *TreeNode{
	if root == nil{
		return nil
	}
	stack := make([]*TreeNode, 0)
	preHead := &TreeNode{}
	p := preHead
	for root != nil || len(stack) > 0 {
		if root != nil{
			stack = append(stack, root)
			root = root.Left
		}else {
			// 出栈处理
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if p != preHead{
				n.Left = p
			}
			p.Right = n
			p = p.Right
			root = n.Right
		}
	}
	// 首尾结点处理
	p.Right = preHead.Right
	if preHead.Right != nil{
		preHead.Right.Left = p
	}
	return preHead.Right
}
