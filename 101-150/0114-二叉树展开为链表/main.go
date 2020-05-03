package _114_二叉树展开为链表

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 类似morris 算法
/*
1.将左子树插入到右子树的地方
2.将原来的右子树接到左子树的最右边节点
3.考虑新的右子树的根节点，一直重复上边的过程，直到新的右子树为 null
*/
/*
			1
		2		3
	4				5
 */
func flatten(root *TreeNode)  {
	if root == nil{
		return
	}
	curr := root
	for curr != nil{
		// 左子树为空，直接找下一个节点
		if curr.Left == nil{
			curr = curr.Right
			continue
		}
		// 寻找左子树最右边的节点
		left := curr.Left
		for left.Right != nil{
			left = left.Right
		}
		// 将原来的右子树接在左子树最右边的节点后
		left.Right = curr.Right
		// 将原来的左子树接在原节点
		curr.Right = curr.Left
		curr.Left = nil // 置空
		// 下一个根节点
		curr = curr.Right
	}
}
