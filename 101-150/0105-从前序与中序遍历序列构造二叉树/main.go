package _105_从前序与中序遍历序列构造二叉树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
	前序: 根 ｜ 左｜ 右
    中序: 左 ｜ 根｜ 右
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{
		Val:   rootVal,
	}
	midRootIndex := findIndexVal(inorder, rootVal)
	if midRootIndex == -1{
		return nil
	}
	leftTreeNodes := midRootIndex
	leftTreePreOrder := preorder[1:leftTreeNodes + 1]
	leftTreeInOrder := inorder[0:midRootIndex]
	rightTreePreOrder := preorder[leftTreeNodes+1:]
	rightTreeInOrder := inorder[midRootIndex+1:]
	root.Left = buildTree(leftTreePreOrder, leftTreeInOrder)
	root.Right = buildTree(rightTreePreOrder, rightTreeInOrder)
	return root
}

func findIndexVal(order []int, target int) int{
	for i, val := range order{
		if val == target{
			return i
		}
	}
	return -1
}