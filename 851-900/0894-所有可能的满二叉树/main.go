package _894_所有可能的满二叉树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
// 递归构建
- 当 N 为偶数时：无法构造满二叉树，返回空数组
- 当 N == 1 时：树只有一个节点，直接返回包含这个节点的数组
- 当完成 N 个节点满二叉树构造时：返回结果数组
*/

func allPossibleFBT(N int) []*TreeNode {
	return helper(N)
}

func helper(N int) []*TreeNode{
	// 1
	if N & 1 == 0{
		return []*TreeNode{}
	}
	// 2
	if N == 1{
		return []*TreeNode{{Val: 0}}
	}
	// 3
	res := make([]*TreeNode, 0)
	for i:=1; i< N; i+=2{
		leftTree := helper(i)
		rightTree := helper(N - 1 - i)
		for _, lNode := range leftTree{
			for _, rNode := range rightTree{
				root := &TreeNode{
					Val:   0,
					Left:  lNode,
					Right: rNode,
				}
				res = append(res, root)
			}
		}
	}
	return res
}
