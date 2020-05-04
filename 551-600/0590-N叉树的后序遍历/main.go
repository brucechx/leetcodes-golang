package _590_N叉树的后序遍历


type Node struct {
	Val int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil{
		return []int{}
	}
	stack := []*Node{root}
	res := make([]int, 0)
	for len(stack) > 0{
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append([]int{node.Val}, res...)
		for _, child := range node.Children{
			stack = append(stack, child)
		}
	}
	return res
}
