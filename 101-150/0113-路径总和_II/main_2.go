package _113_路径总和_II


func MyPathSum(root *TreeNode, target int) [][]int{
	res := make([][]int, 0)
	path := make([]int, 0)
	dfs(root, 0, target, &res, path)
	return res
}

func dfs(root *TreeNode, level, target int, res *[][]int, path []int){
	if root == nil{
		return
	}
	//if level >= len(path){
	path = append(path, root.Val)
	//}else {
	//	path[level] = root.Val
	//}
	target -= root.Val
	if target == 0 && root.Left == nil && root.Right == nil{
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
	}
	for _, node := range []*TreeNode{root.Left, root.Right}{
		dfs(node, level+1, target, res, path)
	}
}
