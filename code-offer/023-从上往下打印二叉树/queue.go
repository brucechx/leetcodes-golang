package _23_从上往下打印二叉树

type Item interface {

}

type Queue struct {
	items []Item
}

func (q *Queue)isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Enqueue(i Item){
	q.items = append(q.items, i)
}

func (q *Queue) Size() int{
	return len(q.items)
}


func (q *Queue) Dequeue() Item{
	if q.isEmpty(){
		return nil
	}
	item := q.items[0]

	q.items = q.items[1:len(q.items)]
	return item
}

func (q *Queue) Pop() Item{
	return q.items[0]
}

