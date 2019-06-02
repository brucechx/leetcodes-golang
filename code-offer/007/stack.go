package _07

type Item interface {

}

type Stack struct {
	Items []Item
}

func initStack() *Stack{
	return &Stack{
		Items:[]Item{},
	}
}

func (s *Stack) Push(t Item){
	s.Items = append(s.Items, t)
}

func (s *Stack) Pop() Item{
	if s.isEmpty(){
		panic("stack is null")
	}
	item := s.Items[len(s.Items) - 1]
	s.Items = s.Items[:len(s.Items) - 1]
	return item
}

func (s *Stack) isEmpty() bool{
	return len(s.Items) == 0
}


