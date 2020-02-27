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
	inputStack := NewStack()
	outputStack := NewStack()

	var res []int
	if root == nil{
		return res
	}

	for root != nil || ! inputStack.IsEmpty(){
		if root != nil{
			inputStack.Push(root)
			outputStack.Push(root)
			root = root.Right
		}else {
			tmp := inputStack.Pop().(*TreeNode)
			root = tmp.Left
		}
	}

	for ! outputStack.IsEmpty(){
		node := outputStack.Pop().(*TreeNode)
		if node == nil{
			continue
		}
		res = append(res, node.Val)
	}
	return res
}


type Stack struct {
	Val []interface{}
}

func NewStack() *Stack{
	return &Stack{Val:make([]interface{}, 0)}
}
func (s *Stack)Push(val interface{}){
	s.Val = append(s.Val, val)
}

func (s *Stack)Pop() interface{}{
	tmp := s.Val[len(s.Val) - 1]
	s.Val = s.Val[:len(s.Val) - 1]
	return tmp
}

func (s *Stack)IsEmpty() bool{
	return len(s.Val) == 0
}