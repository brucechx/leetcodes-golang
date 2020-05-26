package _872_叶子相似的树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var l1, l2 []int
	dfs(root1, &l1)
	dfs(root2, &l2)
	return equal(l1, l2)
}

func equal(l1, l2 []int) bool{
	if len(l1) != len(l2){
		return false
	}
	for i:=0; i<len(l1); i++{
		if l1[i] != l2[i]{
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, res *[]int){
	if root == nil{
		return
	}
	if root.Left == nil && root.Right == nil{
		*res = append(*res, root.Val)
		return
	}
	dfs(root.Left, res)
	dfs(root.Right, res)
}