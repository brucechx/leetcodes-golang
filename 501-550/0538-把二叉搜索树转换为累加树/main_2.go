package _538_把二叉搜索树转换为累加树


// 非递归实现
func convertBST2(root *TreeNode) *TreeNode {
	sum := 0
	st := make([]*TreeNode, 0)
	node := root
	for node != nil || len(st) > 0{
		for node != nil{
			st = append(st, node)
			node = node.Right
		}
		node = st[len(st)-1]
		st = st[:len(st)-1]

		sum += node.Val
		node.Val = sum

		node = node.Left
	}
	return root
}
