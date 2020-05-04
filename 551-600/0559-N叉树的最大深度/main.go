package _559_N叉树的最大深度

type Node struct {
	Val int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil{
		return 0
	}
	max := 0
	for _, child := range root.Children{
		tmp := maxDepth(child)
		if tmp > max{
			max = tmp
		}
	}
	return max + 1
}

