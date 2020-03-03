## [题目描述](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/)
给定一个字符串，找出不含有重复字符的最长子串的长度。

## 示例
- 例1
```text
给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。
给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。
给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列 而不是子串
```

## 思路
- 暴力法

逐个检查所有的子字符串，看它是否不含有重复的字符。

- 滑动窗口

滑动窗口是数组/字符串问题中常用的抽象概念。<br> 
窗口通常是在数组/字符串中由开始和结束索引定义的一系列元素的集合，即 [i, j)[i,j)（左闭，右开）

- 优化的滑动窗口

上述的方法最多需要执行 2n 个步骤。事实上，它可以被进一步优化为仅需要 n 个步骤。我们可以定义字符到索引的映射，而不是使用集合来判断一个字符是否存在。 当我们找到重复的字符时，我们可以立即跳过该窗口。

## 参考
- [滑动窗口技巧](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie/hua-dong-chuang-kou-ji-qiao)