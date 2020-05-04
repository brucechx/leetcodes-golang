package _623_在二叉树中增加一行

type TreeNode struct {
	Val int
	Left *TreeNode
    Right *TreeNode
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1{
		newRoot := &TreeNode{Val:v}
		newRoot.Left = root
		return newRoot
	}
	dfs(root, 1, v, d)
	return root
}

func dfs(root *TreeNode, level, v int, d int){
	if root == nil{
		return
	}
	if level == d - 1{
		left := root.Left
		root.Left = &TreeNode{Val:v}
		root.Left.Left = left
		right := root.Right
		root.Right = &TreeNode{Val:v}
		root.Right.Right = right
	}else {
		dfs(root.Left, level+1, v, d)
		dfs(root.Right, level+1, v, d)
	}

}
