package _538_把二叉搜索树转换为累加树

type TreeNode struct {
	Val int
 	Left *TreeNode
 	Right *TreeNode
}

// 反中序列遍历--递归
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	travel(root, &sum)
	return root
}

func travel(root *TreeNode, sum *int){
	if root == nil{
		return
	}
	travel(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	travel(root.Left, sum)
}
