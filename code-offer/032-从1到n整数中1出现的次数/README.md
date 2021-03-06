[toc]

## 题目
- [题目](https://blog.csdn.net/gatieme/article/details/51292339)

## 题目描述
```text
从1到n整数中1出现的次数
```

## 样例
### 输入
```text
13
```

### 输出
```text
5
1 10 11 12 13 
```

## 思路
### 方法一：暴力方法
最简单直接的方法就是我们循环所有的1~n中的每个number，计算每个number出现的次数

### 方法二：分治递归
任意一个n位数中”1”的个位可以分解为两个n-1位数中”1”的个数的和，最后再加上一个与最高位数相关的常数C
```text
对于n=12，可以拆分为01~09,10~12，即 f(12) = f(10 - 1) + f(12 - 10) + 3,其中3是表示最高位为1的数字个数，这里就是10,11,12；

对于n=132，可以拆分为0~99，100~132，即f(132)=f(100 -1) + f(132 - 100) + 33，33代表最高位为1的数字的个数，这里就是100~132百位数字的1出新了33次

对于232，可以拆分为0~99，100~232，即f(232) = 2*f(100 - 1) + f(32) + 100，因为232大于199，所以它包括了所有最高位为1的数字即100~199，共100个。
```