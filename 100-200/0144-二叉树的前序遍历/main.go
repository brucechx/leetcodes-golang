package main
/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (57.87%)
 * Total Accepted:    16.2K
 * Total Submissions: 27.9K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 前序 遍历。
 *
 * 示例:
 *
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * 输出: [1,2,3]
 *
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
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
	res := []int{}
	if root == nil {
		return res
	}

	cur := root
	for cur != nil {
		if cur.Left == nil {
			res = append(res, cur.Val)
			cur = cur.Right
			continue
		}

		next := cur.Left
		visited := false
		for next.Right != nil {
			if next.Right == cur {
				next.Right = nil
				cur = cur.Right

				visited = true
				break
			}
			next = next.Right
		}

		if !visited {
			next.Right = cur
			res = append(res, cur.Val)
			cur = cur.Left
		}
	}
	return res
}

func preorderTraversal3(root *TreeNode) []int {
	// rightStack 用来暂存右侧节点
	var rightStack []*TreeNode
	var res []int

	for cur := root; cur != nil; {
		res = append(res, cur.Val)

		if cur.Left != nil {
			if cur.Right != nil {
				rightStack = append(rightStack, cur.Right)
			}
			cur = cur.Left
		} else { // cur.Left == nil
			if cur.Right != nil {
				cur = cur.Right
			} else { // cur.Left == cur.Right == nil
				// stack 已空
				// 说明已经完成遍历了
				if len(rightStack) == 0 {
					break
				}
				// 否则
				// 取出最后放入的右侧节点，继续 for 循环
				cur = rightStack[len(rightStack)-1]
				rightStack = rightStack[:len(rightStack)-1]
			}
		}
	}

	return res
}

func preorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for root != nil || len(stack) != 0 {
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			stackSize := len(stack) - 1
			stackPop := stack[stackSize]
			stack = stack[:stackSize]
			root = stackPop.Right
		}
	}
	return res
}

