package _572_另一个树的子树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	for s != nil{
		if dfs(s, t){
			return true
		}
		return isSubtree(s.Left, t) || isSubtree(s.Right, t)
	}
	return false
}

func dfs(s *TreeNode, t *TreeNode) bool{
	if s == nil && t == nil{
		return true
	}
	if s == nil || t == nil{
		return false
	}
	if s.Val != t.Val{
		return false
	}
	return dfs(s.Left, t.Left) && dfs(s.Right, t.Right)
}
