package _106_从中序与后序遍历序列构造二叉树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
 	中序: 左 ｜ 根 ｜ 右
	后序: 左 ｜ 右 ｜ 根

    前  1 2 3
    中  2 1 3
    后  2 3 1
*/

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0{
		return nil
	}
	root := &TreeNode{Val:postorder[len(postorder) - 1]}
	rootValIndex := findIndexVal(inorder, postorder[len(postorder) - 1])
	if rootValIndex == -1{
		return nil
	}
	root.Left = buildTree(inorder[:rootValIndex], postorder[:rootValIndex])
	root.Right = buildTree(inorder[rootValIndex + 1:], postorder[rootValIndex:len(postorder) - 1])
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
