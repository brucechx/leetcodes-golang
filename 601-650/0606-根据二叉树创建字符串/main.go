package _606_根据二叉树创建字符串

import "strconv"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
1. 当前节点右两个孩子，且孩子都不为空；则在两孩子外加括号
2. 当前节点没有孩子，不需要在节点后加括号
3. 当前节点只有左孩子；左孩子加括号，不孩子加括号
4. 当前节点只有右孩子；先空空括号表示左孩子为空，再递归右边孩子
 */

func tree2str(t *TreeNode) string {
	if t == nil{
		return ""
	}
	if t.Left == nil && t.Right == nil{
		return strconv.Itoa(t.Val)
	}
	if t.Right == nil{
		return strconv.Itoa(t.Val) + "("+tree2str(t.Left)+")"
	}
	return strconv.Itoa(t.Val)+"("+tree2str(t.Left)+")"+"("+tree2str(t.Right)+")"
}


