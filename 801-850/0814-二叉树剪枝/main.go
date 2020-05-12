package _814_二叉树剪枝

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if containOne(root){
		return root
	}
	return nil
}

// root 为根子树是否含有1
func containOne(root *TreeNode) bool{
	if root == nil{
		return false
	}
	l := containOne(root.Left)
	r := containOne(root.Right)
	if !l { // 左子树不含，置为空
		root.Left = nil
	}
	if !r { // 同左
		root.Right = nil
	}
	return root.Val == 1 || l || r
}

