[toc]

## 题目
- [题目](https://blog.csdn.net/gatieme/article/details/51407858)

## 题目描述
```text
汇编语言中有一种移位指令叫做循环左移（ROL）

现在有个简单的任务，就是用字符串模拟这个指令的运算结果。 
对于一个给定的字符序列S，请你把其循环左移K位后的序列输出。
```

### 输入 
```text
abcXYZdef  3

```

### 输出
```text
XYZdefabc
```

## 思路
### 方法一： 直接找到旋转后的对应关系
我们可以很明显的发现，将一个字符串循环左移，那么新串和原来串位的对应关系就如下 
* 新串i位置的元素，其实是原来串i+n位置的元素， 
* 当然i+n如果超过了字符串的长度，就会被循环移位到左侧，即(i+n) % str.size() 
* 另外，如果移位n超过了字符串长度，那么只需要移动n%str.size()位即可

### 方法二： 新串是str+str串第n位到底n+str.size()的字串
这个很好理解，由于是循环移位，那么对于abcXYZdef，考虑他的自连接串abcXYZdefabcXYZdef，新的串正好是链接串的一个子串

### 方法三： 通过翻转直线循环移位
通过reverse操作 
一个序列abcdefg，如果向左循环移动n位， 
例如3位，则会编程(defg)(abc)，我们把原序列分成两部分，A=(abc)，B=(defg)，原序列为AB，我们要的结果为BA，则可以这么做：(ArBr)r = ((B)r)r((A)r)r=BA 
假设原序列有n位，循环左移i位的过程如下：

reverse(0,i-1);

reverse(i,n-1);

reverse(1,n-1);

例如原序列：abcdefg，循环左移3位：

abc defg -=> cba defg

cba defg -=> cba gfed

cba gfed -=> defg abc

代码就非常简单了，而且reverse操作非常简单，效率高也不容易出错，要记住一点就是STL中的迭代器是左闭右开区间，所以reverse操作的第二个参数需要往后移动一位
