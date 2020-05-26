package _889_根据前序和后序遍历构造二叉树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 前：根左右； 后： 左右根
func constructFromPrePost(pre []int, post []int) *TreeNode {
	length := len(pre)
	if length == 0{
		return nil
	}
	rootVal := pre[0]
	root := &TreeNode{Val:rootVal}
	if length == 1{
		return root
	}
	// 左子树长度
	l := 0
	for i:=0; i<length; i++{
		if post[i] == pre[1]{
			l = i+1
			break
		}
	}
	root.Left = constructFromPrePost(pre[1:l+1], post[:l])
	root.Right = constructFromPrePost(pre[l+1:], post[l:length-1])
	return root
}
