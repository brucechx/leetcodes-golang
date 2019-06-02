package _21

type MinStack struct {
	Data Stack
	Min Stack
}

func (m *MinStack) Push(i int){
	m.Data.Push(i)
	if m.Min.isEmpty(){
		m.Min.Push(i)
	}else {
		top := m.Min.Top()
		if i < top.(int){
			m.Min.Push(i)
		}else {
			m.Min.Push(top)
		}
	}
}

func (m *MinStack) GetMin() Item{
	m.Data.Pop()
	item := m.Min.Pop()
	return item
}