package _64_数据流之中的中位数

type MaxHeap []int

func (m *MaxHeap) Less(i, j int) bool{
	return (*m)[i] > (*m)[j]
}

func (m *MaxHeap) Swap(i, j int){
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MaxHeap) Len() int{
	return len(*m)
}

func (m *MaxHeap) Push(v interface{}){
	*m = append(*m, v.(int))
}

func (m *MaxHeap) Pop() (v interface{}){
	*m, v = (*m)[:len(*m)-1], (*m)[len(*m) - 1]
	return
}

func (m *MaxHeap) GetTop() interface{}{
	return (*m)[0]
}

