package _60_把二叉树打印成多行

import "fmt"


// 用切片实现队列
type QItem interface {

}

type ArrayQueue struct {
	Items []QItem
}

// 创建队列
func CreateQueue() *ArrayQueue {
	return &ArrayQueue{
		Items: []QItem{},
	}
}

// Enqueue
func (q *ArrayQueue)Enqueue(t QItem)  {
	q.Items = append(q.Items, t)
}

// pop
func (q *ArrayQueue)Pop() QItem {
	if q.isEmpty(){
		panic("queue is empty")
	}
	node := q.Items[0]
	return node
}

// Dequeue
func (q *ArrayQueue)Dequeue() QItem {
	if q.isEmpty(){
		panic("queue is empty")
	}
	node := q.Items[0]
	q.Items = q.Items[1:len(q.Items)]
	return node
}

// isEmpty
func (q *ArrayQueue)isEmpty() bool {
	return len(q.Items) == 0
}

// Size
func (q *ArrayQueue)Size() int {
	return len(q.Items)
}

// 遍历显示队列中的所有元素
func (q *ArrayQueue)printInfo(){
	for _, v := range q.Items{
		fmt.Print(v)
	}
	fmt.Println()
}
