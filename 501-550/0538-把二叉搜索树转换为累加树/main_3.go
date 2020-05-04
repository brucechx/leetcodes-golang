package _538_把二叉搜索树转换为累加树

// 反序中序 Morris 遍历
func convertBST3(root *TreeNode) *TreeNode {
	sum := 0
	node := root
	for node != nil{
		if node.Right == nil{
			sum += node.Val
			node.Val = sum
			node = node.Left
		}else {
			succ := getSuccessor(node)
			if succ.Left == nil{
				succ.Left = node
				node = node.Right
			}else {
				succ.Left = nil
				sum += node.Val
				node.Val = sum
				node = node.Left
			}

		}
	}
	return root
}

func getSuccessor(node *TreeNode) *TreeNode{
	succ := node.Right
	for succ.Left != nil && succ.Left != node{
		succ = succ.Left
	}
	return succ
}
