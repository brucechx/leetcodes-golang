package _25_二叉树中和为某一值的路径

func pathSum(root *Node, sum int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	dfs(&res, root, sum, path)
	return res
}

func dfs(res *[][]int, root *Node, sum int, path []int)  {
	if root == nil{
		return
	}
	path = append(path, root.value)
	if sum - root.value == 0 && root.left == nil && root.right == nil{
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
	}
	dfs(res, root.left, sum-root.value, path)
	dfs(res, root.right, sum-root.value, path)
}
