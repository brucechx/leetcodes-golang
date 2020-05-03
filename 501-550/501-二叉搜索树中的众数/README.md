# [题目描述](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/)
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：
```text
结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树

```
- 例如：
给定 BST [1,null,2,2],
```text
    1
     \
      2
     /
    2
```
返回[2].

提示：如果众数超过1个，不需考虑输出顺序

进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）

## Morris
### 首先明确
在二叉搜索树中，如果一个结点有前驱结点，那么前驱结点的右指针只有两种情况（可以自己找两个二叉搜索树验证一下，此处不再证明）
- 是空的
- 是这个结点本身（即前驱是它的父结点）

所以我们可以把前驱结点的右指针这一特性利用起来，从而降低空间复杂度

### Morris遍历算法的步骤如下：

检查当前结点的左孩子:
- 如果当前结点的左孩子为空，说明要不没有前驱，要不前驱是它的父结点，所以进行检查，然后进入右孩子。
- 如果当前结点的左孩子不为空，说明左子树里肯定有它的前驱，那就找到这个前驱
    + 如果前驱结点的右孩子是空，说明还没检查过左子树，那么把前驱结点的右孩子指向当前结点，然后进入当前结点的左孩子。
    + 如果当前结点的前驱结点其右孩子指向了它本身，说明左子树已被检查过，就直接进行检查，然后把前驱结点的右孩子设置为空，恢复原树，再进入右孩子。

### Morris代码框架
```text
while(cur){
    if(cur->left == NULL){// 左子树为空时，直接比较，然后进入右子树
        /*************
        /*  进行你的操作
        *************/
        pre = cur;
        cur = cur->right;
    }else{// 进入左子树
        /*************
        /* 找cur的前驱结点，找到后分两种情况
        /*   1、cur的左子结点没有右子结点，那cur的左子结点就是前驱
        /*   2、cur的左子结点有右子结点，就一路向右下，走到底就是cur的前驱
        *************/
        TreeNode* preceesor = cur->left;
        while(preceesor->right && preceesor->right != cur){
            preceesor = preceesor->right;
        }

        // 前驱已经指向自己了，直接比较，然后进入右子树
        if(preceesor->right == cur){
            /*************
            /*  进行你的操作
            *************/
            preceesor->right = NULL; // 断开连接，恢复原树
            pre = cur;
            cur = cur->right;
        }else{ // 前驱还没有指向自己，说明左边还没有遍历，将前驱的右指针指向自己，后进入前驱
            preceesor->right = cur;
            cur = cur->left;
        }
    }
}
```

## 参考
- [link](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/solution/lian-di-gui-du-bu-yong-de-chang-shu-kong-jian-fu-z/)
