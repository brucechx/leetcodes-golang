package _37_两个链表的第一个公共结点

/*
	入栈法
*/

type Node struct {
	Val int
	Next *Node
}


type Stack struct {
	Val []Node
}

func NewStack() *Stack{
	return &Stack{}
}

func (s *Stack) Push(n Node){
	s.Val = append(s.Val, n)
}

func (s *Stack) Pop() Node{
	tmp := s.Val[len(s.Val) - 1]
	s.Val = s.Val[: len(s.Val) - 1]
	return tmp
}

func (s *Stack) IsEmpty() bool{
	return len(s.Val) == 0
}


func FindFirstCommonNode(head1, head2 *Node) int{
	stack1 := NewStack()
	stack2 := NewStack()

	for node:=head1; node!=nil; node=node.Next{
		stack1.Push(*node)
	}

	for node:=head2; node!=nil; node=node.Next{
		stack2.Push(*node)
	}

	for !stack1.IsEmpty() || !stack2.IsEmpty(){
		s1 := stack1.Pop()
		s2 := stack2.Pop()
		if s1.Val != s2.Val{
			return s1.Next.Val
		}
	}
	return -1
}



