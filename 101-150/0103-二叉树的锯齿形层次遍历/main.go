package _103_二叉树的锯齿形层次遍历

func PrintTreeNew(root *TreeNode) [][]int{
	res := make([][]int, 0)
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, level int, res *[][]int){
	if root == nil{
		return
	}
	if len(*res) == level{
		*res = append(*res, []int{})
	}
	if level % 2 != 0{
		(*res)[level] = append([]int{root.Val}, (*res)[level]...)
	}else {
		(*res)[level] = append((*res)[level], root.Val)
	}
	for _, node := range []*TreeNode{root.Left, root.Right}{
		dfs(node, level+1, res)
	}
}



