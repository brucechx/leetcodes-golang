[toc]

## 题目 
- [题目](https://leetcode-cn.com/problems/zigzag-conversion/)

## 描述

将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);

## 示例 
**示例1**:
```text
输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
```

**示例2**:
```text
输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

```

## 思路
### 方法一：
输入"ABCDEFGHIJKLMNOPQRSTUVWXYZ"和参数5后，得到答案"AGMSYBFHLNRTXZCEIKOQUWDJPV"， 按照题目的摆放方法，可得：
```text
A   I   Q   Y
B  HJ  PR  XZ
C G K O S W
DF  LN  TV
E   M   U
```

可以看到，各行字符在原字符串中的索引号为
```text
0行，0, 8, 16, 24
1行，1, 7,9, 15,17, 23,25
2行，2, 6, 10, 14, 18, 22
3行，3,5, 11,13, 19,21
4行，4, 12, 20
```

令p=numRows×2-2，可以总结出以下规律
```text
0行， 0×p，1×p，...
r行， r，1×p-r，1×p+r，2×p-r，2×p+r，...
最后一行， numRow-1, numRow-1+1×p，numRow-1+2×p，...
```
只需编程依次处理各行即可。