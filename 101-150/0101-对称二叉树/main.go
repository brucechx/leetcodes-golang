package _101_对称二叉树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

/*
 	递归
*/
func isMirror(t1, t2 *TreeNode) bool{
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil{
		return false
	}
	return t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil{
		return true
	}
	currentNodes := make([]*TreeNode, 0)
	currentNodes = append(currentNodes, root)
	for len(currentNodes) != 0 {
		nextSeqNodes := make([]*TreeNode, 0)
		diffNodes := make([]*TreeNode, 0)
		for _, node := range currentNodes{
			diffNodes = append(diffNodes, node.Left)
			diffNodes = append(diffNodes, node.Right)
			if node.Left != nil{
				nextSeqNodes = append(nextSeqNodes, node.Left)
			}
			if node.Right != nil{
				nextSeqNodes = append(nextSeqNodes, node.Right)
			}
		}
		if ! isSameNodes(diffNodes){
			return false
		}
		currentNodes = nextSeqNodes
	}
	return true
}

func isSameNodes(nodes []*TreeNode) bool{
	if len(nodes) == 1{
		return true
	}
	i := 0
	j := len(nodes) - 1
	mid := len(nodes) / 2
	for i < mid{
		if ! isSameTreeNode(nodes[i], nodes[j]){
			return false
		}
		i++
		j--
	}
	return true
}

func isSameTreeNode(p, q *TreeNode) bool{
	if p == nil && q == nil{
		return true
	}
	if p == nil && q != nil{
		return false
	}
	if p != nil && q == nil{
		return false
	}
	return p.Val == q.Val
}