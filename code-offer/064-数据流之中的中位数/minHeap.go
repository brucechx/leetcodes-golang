package _64_数据流之中的中位数

type MinHeap []int

func (m *MinHeap) Less(i, j int) bool{
	return (*m)[i] < (*m)[j]
}

func (m *MinHeap) Swap(i, j int){
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MinHeap) Len() int{
	return len(*m)
}

func (m *MinHeap) Push(v interface{}){
	*m = append(*m, v.(int))
}

func (m *MinHeap) Pop() (v interface{}){
	*m, v = (*m)[:len(*m)-1], (*m)[len(*m) - 1]
	return
}

func (m *MinHeap) GetTop() interface{}{
	return (*m)[0]
}
