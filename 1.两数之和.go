package leetcode

import "errors"

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 *
 * https://leetcode-cn.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (45.14%)
 * Total Accepted:    273.8K
 * Total Submissions: 606.7K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 *
 * 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
 *
 * 示例:
 *
 * 给定 nums = [2, 7, 11, 15], target = 9
 *
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 *
 *
 */

/*
	解题思路：
	方法三：一遍哈希表
	在进行迭代并将元素插入到表中的同时，我们还会回过头来检查表中是否已经存在当前元素所对应的目标元素。
	如果它存在，那我们已经找到了对应解，并立即将其返回。
*/
func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	for i, v := range nums {
		t := target - v
		if _, ok := tmp[t]; ok {
			return []int{tmp[t], i}
		}
		tmp[v] = i
	}
	panic(errors.New("errors"))
}
