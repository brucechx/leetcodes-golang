package _145_二叉树的后序遍历

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
 	后序遍历: 左右根
    出栈： 根右左
*/

func postorderTraversal(root *TreeNode) []int {
	inStack := NewStack()
	outStack := NewStack()
	for root != nil || ! inStack.isEmpty(){
		if root != nil{
			inStack.Put(root)
			outStack.Put(root)
			root = root.Right
		}else{
			root = inStack.Pop()
			root = root.Left
		}
	}
	res := make([]int, 0)
	for ! outStack.isEmpty(){
		node := outStack.Pop()
		res = append(res, node.Val)
	}
	return res
}

type Stack struct{
	data []*TreeNode
}

func NewStack() *Stack{
	return &Stack{
		data: make([]*TreeNode, 0),
	}
}

func (s *Stack) Put(node *TreeNode){
	s.data = append(s.data, node)
}

func (s *Stack) Pop() *TreeNode{
	if len(s.data) == 0{
		return nil
	}
	val := s.data[len(s.data) - 1]
	s.data = s.data[: len(s.data) - 1]
	return val
}

func (s *Stack) isEmpty() bool{
	return len(s.data) == 0
}