package _56_链表中环的入口结点

//链表结点
type Node struct{
	data interface{}
	next *Node
}

//链表
type LinkList struct{
	head *Node
	tail *Node
	size int
}

//新建空链表，即创建Node指针head，用来指向链表第一个结点，初始为空
func NewLinkList() *LinkList{
	return &LinkList{}
}

//在链表尾部添加数据
func(l *LinkList) append(e interface{}){
	//为数据创建新结点
	newNode:=Node{}
	newNode.data=e
	newNode.next=nil

	if l.size==0 {
		l.head=&newNode
		l.tail=&newNode
	}else{
		l.tail.next=&newNode
		l.tail=&newNode
	}
	l.size++
}

//获取含有指定数据的第一个结点
func(l *LinkList) getNode(e interface{}) *Node{
	var p *Node=l.head
	for p!=nil{
		//找到该数据所在结点
		if e==p.data{
			return p
		}else{
			p=p.next
		}
	}
	return nil
}

// 如果链表有环，输出入环点
/*
	* 判断链表是否有环
 	* 判断方法是设置两个指针最初均指向头结点，然后fast每次走2步，slow每次走1步，
 	* 如果链表没有环，则fast指针肯定先指向表尾的null
 	* 如果有环，则fast和slow肯定会相遇。然后第一次相遇后我们将fast指针重新指向头结点，
 	* 然后fast和slow每次都走一步直到第二次相遇，那么第二次相遇的节点即为入环的节点
 */
func (l *LinkList) CircleNode(n *Node) *Node{
	slow := n
	fast := n
	for fast != nil && fast.next.next != nil{
		slow = slow.next
		fast = fast.next.next
		if slow == fast{
			break
		}
	}

	p3 := slow
	p4 := n
	for {
		p3 = p3.next
		p4 = p4.next
		if p3.data == p4.data{
			return p3
		}
	}
	return nil
}

// 判断链表是否有环，双指针
func (l *LinkList) circle(n *Node) bool{
	p1 := n
	p2 := n
	for p2 != nil && p2.next != nil{
		p1 = p1.next
		p2 = p2.next.next
		if p1 == p2{
			return true
		}
	}
	return false
}