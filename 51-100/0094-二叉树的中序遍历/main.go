package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func InorderTraversalWithRec(root *TreeNode) []int{
	var res []int
	inorderTraversalWithRec(root, &res)
	return res
}

func InorderTraversalWithStack(root *TreeNode) []int{
	return inorderTraversalWithStack(root)
}

// 递归实现
// 如果传 res []int 为值传递
func inorderTraversalWithRec(node *TreeNode, res *[]int) {
	// 左根右
	if node == nil{
		return
	}
	inorderTraversalWithRec(node.Left, res)
	*res = append(*res, node.Val)
	inorderTraversalWithRec(node.Right, res)
}

// 非递归实现-基于栈
func inorderTraversalWithStack(node *TreeNode) []int{
	s := NewStack()
	var res []int

	for node != nil || ! s.isEmpty(){
		if node != nil{
			s.Push(node)
			node = node.Left
		}else{
			tmp := s.Pop()
			res = append(res, tmp.Val)
			node = tmp.Right
		}
	}
	return res
}


type Stack struct {
	Val []*TreeNode
}

func NewStack() *Stack{
	return &Stack{Val:[]*TreeNode{}}
}

func (s *Stack) Push(v *TreeNode){
	s.Val = append(s.Val, v)
}

func (s *Stack) Pop() *TreeNode{
	sSize := len(s.Val) - 1
	tmp := s.Val[sSize]
	s.Val = s.Val[:sSize]
	return tmp
}

func (s *Stack) isEmpty() bool{
	return len(s.Val) == 0
}

// 基于栈
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			stackSize := len(stack) - 1
			stackPop := stack[stackSize]
			stack = stack[:stackSize]

			res = append(res, stackPop.Val)
			root = stackPop.Right
		}
	}

	return res
}

