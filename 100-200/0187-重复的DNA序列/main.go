package main
import "sort"

/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 *
 * https://leetcode-cn.com/problems/repeated-dna-sequences/description/
 *
 * algorithms
 * Medium (41.80%)
 * Total Accepted:    2.2K
 * Total Submissions: 5.3K
 * Testcase Example:  '"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"'
 *
 * 所有 DNA 由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA
 * 中的重复序列有时会对研究非常有帮助。
 *
 * 编写一个函数来查找 DNA 分子中所有出现超多一次的10个字母长的序列（子串）。
 *
 * 示例:
 *
 * 输入: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
 *
 * 输出: ["AAAAACCCCC", "CCCCCAAAAA"]
 *
 */

func findRepeatedDnaSequences(s string) []string {
	var res []string
	if len(s) <= 10 {
		return nil
	}

	str := ""
	// rec 记录各个 subString 出现的次数
	rec := make(map[string]int, len(s)-9)
	for i := 0; i+10 <= len(s); i++ {
		str = s[i : i+10]
		if v := rec[str]; v == 1 {
			// 把已经出现过一次的 subString 存放入 res
			res = append(res, str)
		}
		rec[str]++
	}

	sort.Strings(res)

	return res
}

