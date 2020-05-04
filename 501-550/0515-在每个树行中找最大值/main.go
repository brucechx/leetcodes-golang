package _515_在每个树行中找最大值

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, level int, res *[]int){
	if root == nil{
		return
	}
	if level >= len(*res){
		*res = append(*res, root.Val)
	}else {
		if root.Val > (*res)[level]{
			(*res)[level] = root.Val
		}
	}
	dfs(root.Left, level+1, res)
	dfs(root.Right, level+1, res)

}
