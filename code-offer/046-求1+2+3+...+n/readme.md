[toc]

## 题目
- [题目](https://blog.csdn.net/gatieme/article/details/51485713)

## 题目描述
```text
求1+2+3+…+n，
要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
```

### 输入 
```text
3

5
```

### 输出
```text
6

15
```

## 思路
### 方法一： 递归+短路判断终止

### 方法二： 位运算计算
我们知道a*b运算在计算机内部其实是通过移位和加法来完成的
```text
int Multi(int a, int b)
{
    int res = 0;
    while(a != 0)
    {
        if((a & 1) != 0)
        {
            res += b;
        }
        a >>= 1;
        b <<= 1;

    }

    return res;
}
```
那么我们的结果就是
```text
Multi(n, n + 1) >>1 
但是我们的乘法函数中仍然有循环和判断, 因此我们需要进一步优化,
```

还是上面的策略， 
* 循环用递归来代替 
* 分支用短路运算来代替

那么我们的乘法运算就成为
```text
int MultiRecursion(int a, int b)
{
    static int res = 0;
    a && MultiRecursion(a >> 1, b << 1);    //  递归的进行运算
    (a & 1) && (res += b);                  //  短路

    return res;
}
```
