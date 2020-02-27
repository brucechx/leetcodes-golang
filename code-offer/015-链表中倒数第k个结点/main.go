package _15_链表中倒数第k个结点

type Node struct {
	value int
	next *Node
}

// LastK 获取倒数第 K 个节点
// 相当于第 length - K + 1 个节点
// 采用双指针法，减少第二次遍历
// 第一个指针right先向前走K步，然后left和right一起走，
// 此时两个指针差别K步，那么当right走到链表尾部的时候，
// left指向的就是倒数第K个节点
func LastK(head *Node, k int) *Node{
	// 场景一
	if head == nil || k <= 0{
		return nil
	}

	leftNode := head
	rightNode := head
	for i:=0; i<k; i++{
		// 链表长度小于k
		if rightNode == nil{
			return nil
		}
		rightNode = rightNode.next
	}

	// rightNode 到达了第一个空节点
	// 不能写在 rightNode.next 后面
	// 因为当链表长度为 k 时，rightNode 为 nil
	// 造成 nil pointer dereference
	for {
		if rightNode == nil{
			return leftNode
		}
		leftNode = leftNode.next
		rightNode = rightNode.next
	}
}
