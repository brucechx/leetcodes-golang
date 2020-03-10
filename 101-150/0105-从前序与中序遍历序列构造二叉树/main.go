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
	if len(inorder) == 0{
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	rootIndex := findIndex(inorder, preorder[0])
	if rootIndex == -1{
		return nil
	}
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}

func findIndex(order []int, target int) int{
	for i, v := range order{
		if v == target{
			return i
		}
	}
	return -1
}