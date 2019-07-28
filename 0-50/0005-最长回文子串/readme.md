[toc]

## 题目
- [题目](
https://leetcode-cn.com/problems/longest-palindromic-substring/description/)
- [解说](https://leetcode-cn.com/articles/longest-palindromic-substring/)

## desc

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为1000。

**示例**:
```text
输入: "babad"
输出: "bab"
注意: "aba"也是一个有效答案。
```

## 思路
### 方法一：最长公共子串
### 方法二：暴力法
很明显，暴力法将选出所有子字符串可能的开始和结束位置，并检验它是不是回文。
### 方法三：动态规划
如果 p(i, j) 是回文，则 p(i+1, j-1)且si==sj

基本示例:
1个字符  p(i, i) = true
2个字符  p(i, i+1) = (si == si+1)
### 方法四：中心扩展算法
回文中心的两侧互为镜像。因此，回文可以从他的中心展开，并且只有2n-1个这样的中心(一个元素为中心的情况有n个，两个元素为中心的情况有n-1个)
### 方法五：Manacher 算法