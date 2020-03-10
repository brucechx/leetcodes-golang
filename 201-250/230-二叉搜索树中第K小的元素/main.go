package _30_二叉搜索树中第K小的元素

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil{
		return -1
	}
	stack := NewStack()
	i := 1
	for root != nil || !stack.IsEmpty(){
		if root != nil{
			stack.Push(root)
			root = root.Left
		}else {
			tmp := stack.Pop()
			if i == k{
				return tmp.Val
			}
			i++
			root = tmp.Right
		}
	}
	return -1
}

type Stack struct {
	data []*TreeNode
}

func NewStack() *Stack{
	return &Stack{data:make([]*TreeNode, 0)}
}

func (s *Stack)Push(v *TreeNode){
	s.data = append(s.data, v)
}

func (s *Stack)Pop() *TreeNode{
	v := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *Stack)IsEmpty() bool{
	return len(s.data) == 0
}