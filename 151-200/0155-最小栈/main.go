package _155_最小栈

/*
	两种实现方式：
	1。 两个栈一样大小
	2。 最小栈只保留最小值
*/

type MinStack struct {
	originData []int
	minData []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		originData: make([]int, 0),
		minData:    make([]int, 0),
	}
}


func (t *MinStack) Push(x int)  {
	t.originData = append(t.originData, x)
	if len(t.minData) == 0{
		t.minData = append(t.minData, x)
		return
	}
	minVal := t.minData[len(t.minData) - 1]
	if x < minVal{
		t.minData = append(t.minData, x)
	}else {
		t.minData = append(t.minData, minVal)
	}
}


func (t *MinStack) Pop()  {
	if len(t.minData) == 0{
		return
	}
	t.minData = t.minData[:len(t.minData)-1]
	t.originData = t.originData[:len(t.originData)-1]
}


func (t *MinStack) Top() int {
	if len(t.originData) == 0{
		return -1
	}
	val := t.originData[len(t.originData)-1]
	return val
}


func (t *MinStack) GetMin() int {
	if len(t.minData) == 0{
		return -1
	}
	val := t.minData[len(t.minData)-1]
	return val
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
