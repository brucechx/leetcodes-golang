package _865_有所有最深结点的最小子树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 一次深度遍历；返回最深节点最小子树的根节点以及最深距离
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	node, _ := dfs(root)
	return node
}

func dfs(node *TreeNode) (subNode *TreeNode, dist int) {
	if node == nil{
		return
	}
	l, lLen := dfs(node.Left)
	r, rLen := dfs(node.Right)
	if lLen > rLen{
		return l, lLen + 1
	}else if lLen < rLen{
		return r, rLen + 1
	}
	return node, lLen + 1
}
