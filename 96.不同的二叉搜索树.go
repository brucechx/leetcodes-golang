/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (58.14%)
 * Total Accepted:    5.8K
 * Total Submissions: 10K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
 *
 * 示例:
 *
 * 输入: 3
 * 输出: 5
 * 解释:
 * 给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
 */
func numTrees(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 || n == 2 {
		return n
	}

	if n == 3 {
		return 5
	}

	res := 0
	// 左右对称，所以只做一半
	for i := 1; i <= n/2; i++ {
		res += numTrees(i-1) * numTrees(n-i)
	}
	res *= 2
	// 处理 n 为奇数的情况
	if n%2 == 1 {
		temp := numTrees(n / 2)
		res += temp * temp
	}

	return res
}

