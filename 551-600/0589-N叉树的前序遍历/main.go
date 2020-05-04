package _589_N叉树的前序遍历

type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil{
		return []int{}
	}
	res := make([]int, 0)
	stack := []*Node{root}
	for len(stack) > 0{
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res, node.Val)
		for i:=len(node.Children)-1; i>=0; i--{
			stack = append(stack, node.Children[i])
		}
	}
	return res
}
