# [题目描述](https://leetcode-cn.com/problems/word-ladder/)
给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：

每次转换只能改变一个字母。
转换过程中的中间单词必须是字典中的单词。
说明:

如果不存在这样的转换序列，返回 0。
所有单词具有相同的长度。
所有单词只由小写字母组成。
字典中不存在重复的单词。
你可以假设 beginWord 和 endWord 是非空的，且二者不相同。

```text
示例 1:
输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出: 5

解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
     返回它的长度 5。

示例 2:
输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

输出: 0

解释: endWord "cog" 不在字典中，所以无法进行转换。

```

## 思路
- BFS
```text
BFS:
	q.push(start) // 初始节点放入队列
step = 0 // 初始化步长为0

while not q.empty(): // 当前节点不为空
    ++step // 扩展一层节点
    size = q.size() // 当前这一层节点的个数
    while size-- > 0: // 当前这一层节点数量大于0
        node = q.pop() // 不断出队
        new_nodes = expand(node) // 扩展下一层的节点
        if goal in new_nodes: return step + 1 // 找到结果step+1返回步长
        q.append(new_nodes) // 如果没有找到结果，把下一层的所有节点加入队列

return NOT_FOUND // 队列全部结束，没有找到结果
```
- 双向BFS
```text
s1.insert(start) // 方便使用HashSet
s2.insert(end) // start和end放入两个set
step = 0 // 初始化步长为0

while not (s1.empty() || s2.empty()): // 当两个set都不为空，循环执行
    ++step // 步长+1
    swap*(s1, s2) // 交替从左端扩展和从右端扩展
    s = {} // 定义新的空集合
    for node in s1: // 当前需要扩展的这一层节点进行遍历
        new_nodes = expand(node) // 扩展下一层节点
        if any(new_nodes) in s2: return step + 1 // 新的节点在s2集合中，返回step+1，找到路径
        s.append(new_nodes) // 如果没有找到结果，把下一层的所有节点加入队列
    s1 = s // 临时的集合赋值给s1，即把已经走的路径存储起来

return NOT_FOUND // 队列全部结束，没有找到结果
```
