package _65_滑动窗口的最大值

type Deque struct {
	data []int // idx
}

func NewDeque() *Deque{
	return &Deque{}
}

func (d *Deque) isEmpty() bool{
	return len(d.data) == 0
}

func (d *Deque) getLast() int{
	return d.data[len(d.data) - 1]
}

func (d *Deque) removeLast() {
	d.data = d.data[:len(d.data) - 1]
}

func (d *Deque) removeFirst() {
	d.data = d.data[1:]
}

func (d *Deque) addLast(i int){
	d.data = append(d.data, i)
}

func (d *Deque) getFirst() int{
	return d.data[0]
}

func MaxInWindows(data []int, size int) (maxInWindows []int){
	// 边界
	if len(data) < 1 || size < 1 {
		return
	}
	dq := NewDeque()
	// 窗口还没有被填满时，找最大值的索引
	for i:=0; i< size && i< len(data); i++{
		// 如果索引对应的值比之前存储的索引值对应的值大或者相等，就删除之前存储的值
		for !dq.isEmpty() && data[i] >= data[dq.getLast()]{
			dq.removeLast()
		}
		// 添加索引
		dq.addLast(i)
	}
	// 窗口已经被填满了
	for i:=size; i< len(data); i++{
		// 第一个窗口的最大值保存
		maxInWindows = append(maxInWindows, data[dq.getFirst()])
		// 如果索引对应的值比之前存储的索引值对应的值大或者相等，就删除之前存储的值
		for !dq.isEmpty() && data[i] >= data[dq.getLast()]{
			dq.removeLast()
		}
		// 删除已经滑出窗口的数据对应的下标
		for !dq.isEmpty() && dq.getFirst() <= i - size{
			dq.removeFirst()
		}
		// 可能的最大的下标索引入队
		dq.addLast(i)
	}
	// 最后一个窗口最大值入队
	maxInWindows = append(maxInWindows, data[dq.getFirst()])
	return

}
