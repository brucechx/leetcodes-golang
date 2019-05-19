/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (60.52%)
 * Total Accepted:    15.1K
 * Total Submissions: 25K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 * 示例 1:
 *
 * 输入: nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出: [2]
 *
 *
 * 示例 2:
 *
 * 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出: [9,4]
 *
 * 说明:
 *
 *
 * 输出结果中的每个元素一定是唯一的。
 * 我们可以不考虑输出结果的顺序。
 *
 *
 */

func intersection(nums1 []int, nums2 []int) []int {

	var res []int
	if len(nums1) < 0 || len(nums2) < 0 {
		return res
	}
	m1 := getInts(nums1)
	m2 := getInts(nums2)

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}

	for i := range m1 {
		if m2[i] {
			res = append(res, i)
		}
	}
	return res
}

func getInts(nums []int) map[int]bool {
	intMap := make(map[int]bool, len(nums))
	for _, v := range nums {
		intMap[v] = true
	}
	return intMap
}

