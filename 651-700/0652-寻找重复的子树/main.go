package _652_寻找重复的子树

import "strconv"

type TreeNode struct {
	Val int
 	Left *TreeNode
 	Right *TreeNode
}

// 递归序列化 + hash

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	count := make(map[string]int)
	res := make([]*TreeNode, 0)
	helper(root, count, &res)
	return res
}

func helper(root *TreeNode, count map[string]int, res *[]*TreeNode) string{
	if root == nil{
		return "*"
	}
	l := helper(root.Left, count, res)
	r := helper(root.Right, count, res)

	key := strconv.Itoa(root.Val) + l + r
	count[key]++
	if count[key] == 2{
		*res = append(*res, root)
	}
	return key
}
