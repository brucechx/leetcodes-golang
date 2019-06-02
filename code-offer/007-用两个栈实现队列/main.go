package _07_用两个栈实现队列

type StackQueue struct {
	stack1 *Stack
	stack2 *Stack
}

func initStackQueue() *StackQueue{
	return &StackQueue{
		stack1:initStack(),
		stack2:initStack(),
	}
}

// 入队列
func (sq *StackQueue) Enqueue(i Item){
	sq.stack1.Push(i)
}

// 出队列
func (sq *StackQueue) Dequeue() Item{
	if sq.stack2.isEmpty(){
		for ! sq.stack1.isEmpty(){
			tmp := sq.stack1.Pop()
			sq.stack2.Push(tmp)
		}
	}
	return sq.stack2.Pop()
}