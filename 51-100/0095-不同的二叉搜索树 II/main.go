package main

import "fmt"

/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (51.19%)
 * Total Accepted:    3.8K
 * Total Submissions: 7.4K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * 解释:
 * 以上的输出对应以下 5 种不同结构的二叉搜索树：
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	var res []*TreeNode

	if n <= 0 {
		return res
	}

	// 所有符合条件的二叉树的 inorder 是一样的。
	in := make([]int, n)
	for i := range in {
		in[i] = i + 1
	}

	// 利用 inorder 生成所有可能的 preorder
	pres := getPres(in)

	// 利用 preorder 和 inorder 生成 二叉树
	for _, pre := range pres {
		temp := preIn2Tree(pre, in)
		res = append(res, temp)
	}

	return res
}

func getPres(in []int) [][]int {
	size := len(in)
	if size <= 1 {
		return [][]int{in}
	}

	if size == 2 {
		return [][]int{
			[]int{in[1], in[0]},
			[]int{in[0], in[1]},
		}
	}

	res := [][]int{}
	for i := 0; i < size; i++ {
		// 以 in[i] 为 root
		// 获取 in[i] 左侧的 preorder
		ls := getPres(in[:i])
		// 获取 in[i] 右侧的 preorder
		rs := getPres(in[i+1:])
		for _, l := range ls {
			for _, r := range rs {
				temp := make([]int, 1, size)
				// in[i] 为 root，所以，应该在 0 位
				temp[0] = in[i]
				temp = append(temp, l...)
				temp = append(temp, r...)
				// 汇总结果
				res = append(res, temp)
			}
		}
	}

	return res
}

func preIn2Tree(pre, in []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = preIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = preIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	msg := fmt.Sprintf("%d 不存在于 %v 中", val, nums)
	panic(msg)
}