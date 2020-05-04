package _501_二叉搜索树中的众数

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

// Morris
/*
1. 中序遍历
2. 如果当前节点与前一个节点相等；count++；
3. 如果count > maxCount 则覆盖众数数组
4. 如果相等，则追加，否则忽略
*/

func findMode(root *TreeNode) []int {
	var pre *TreeNode
	curr := root
	count, maxCount := 0, 0
	res := make([]int, 0)

	for curr != nil{
		if curr.Left == nil{ // 如果左子树为空，直接比较，然后进入右子树
			if pre != nil && pre.Val == curr.Val{
				count++
			}else {
				count = 1
			}
			if count > maxCount{
				maxCount = count
				res = []int{curr.Val}
			}else if count == maxCount{
				res = append(res, curr.Val)
			}
			pre = curr
			curr = curr.Right
		}else { // 左子树不为空，进行morris 遍历
			// 找到cur的前驱结点，分两种情况
			// 1、cur的左子结点没有右子结点，那cur的左子结点就是前驱
			// 2、cur的左子结点有右子结点，就一路向右下，走到底就是cur的前驱
			processor := curr.Left
			for processor.Right != nil && processor.Right != curr{
				processor = processor.Right
			}
			// 前驱已经指向自己了，直接比较，然后进入右子树
			if processor.Right == curr{
				// 比较运算
				if pre != nil && pre.Val == curr.Val{
					count++
				}else {
					count = 1
				}
				if count > maxCount{
					maxCount = count
					res = []int{curr.Val}
				}else if count == maxCount{
					res = append(res, curr.Val)
				}
				// morris 遍历
				processor.Right = nil
				pre = curr
				curr = curr.Right
			}else {
				processor.Right = curr
				curr = curr.Left
			}
		}
	}
	return res
}
