package _429_N叉树的层序遍历

type Node struct {
	Val int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	dfs(root, 0, &res)
	return res
}

func dfs(root *Node, level int, res *[][]int){
	if root == nil{
		return
	}
	if level == len(*res){
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], root.Val)
	for _, node := range root.Children{
		dfs(node, level+1, res)
	}
}
