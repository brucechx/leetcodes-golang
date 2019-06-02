package _05_从尾到头打印链表

type Node struct {
	value interface{}
	next *Node
}

type SingleList struct {
	head *Node
	tail *Node
	size int
}

func newSingeList() *SingleList{
	return &SingleList{
		head:nil,
		tail:nil,
		size:0,
	}
}

func (s *SingleList) append(v interface{}){
	node := &Node{
		value:v,
		next:nil,
	}
	if s.size == 0{
		s.head = node
		s.tail = node
	}else {
		s.tail.next = node
		s.tail = node
	}
	s.size ++
}

func (s *SingleList)printReverse() []interface{}{
	var tmp []interface{}
	node := s.head
	for node != nil{
		// 入栈
		tmp = append(tmp, node.value)
		node = node.next
	}
	var result []interface{}
	for i:=len(tmp)-1; i >=0; i-- {
		// 出栈
		result = append(result, tmp[i])
	}
	return result
}
