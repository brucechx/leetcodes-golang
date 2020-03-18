package _146_LRU缓存机制

type Node struct {
	key int
	val int
	pre *Node
	next *Node
}

type LRUCache struct {
	data map[int]*Node
	head *Node
	tail *Node
	capacity int
	count int
}


func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		data:     make(map[int]*Node),
		head:     head,
		tail:     tail,
		capacity: capacity,
		count:    0,
	}
}

func (l *LRUCache) Get(key int) int {
	if v, ok := l.data[key]; ok{
		// 先删除结点，然后置于缓存首结点
		l.detachNode(v)
		l.insertFront(v)
		return v.val
	}
	return -1
}


func (l *LRUCache) Put(key int, value int)  {
	if v, ok := l.data[key]; !ok{
		node := &Node{key:key, val:value}
		if l.capacity == l.count{
			l.delTail()
		}
		l.data[key] = node
		l.insertFront(node)
		l.count++
	}else {
		l.detachNode(v)
		v.val = value // 值可以更新；如果不需更新，则无此步骤
		l.insertFront(v)
	}
}

func (l *LRUCache) detachNode(node *Node){
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (l *LRUCache) insertFront(node *Node){
	node.next = l.head.next
	l.head.next.pre = node
	l.head.next = node
	node.pre = l.head
}

func (l *LRUCache) delTail(){
	tmp := l.tail.pre
	tmp.pre.next = l.tail
	l.tail.pre = tmp.pre
	tmp.next = nil
	tmp.pre = nil
	l.count --
	delete(l.data, tmp.key)
}
