package _21_包含min函数的栈

type Item interface {

}

type Stack struct {
	items []Item
}

func (s *Stack) isEmpty() bool{
	return len(s.items) == 0
}

func (s *Stack) Push(i Item){
	s.items = append(s.items, i)
}

func (s *Stack) Top() Item{
	if s.isEmpty(){
		return nil
	}
	return s.items[len(s.items) - 1]
}

func (s *Stack) Pop() Item{
	if s.isEmpty(){
		return nil
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

