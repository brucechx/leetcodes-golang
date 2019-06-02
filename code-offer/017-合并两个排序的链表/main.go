package _17_合并两个排序的链表

type Node struct {
	val int
	next *Node
}

func SortedMerge(head1, head2 *Node) *Node{
	if head1 == nil{
		return head2
	}
	if head2 == nil{
		return head1
	}

	var mergeHead *Node
	if head1.val < head2.val{
		mergeHead = head1
		mergeHead.next = SortedMerge(head1.next, head2)
	}else {
		mergeHead = head2
		mergeHead.next = SortedMerge(head1, head2.next)
	}
	return mergeHead
}


func SortedMerge2(l1 *Node, l2 *Node) (result *Node) {
	l1Index := l1
	l2Index := l2
	var current *Node = nil

	for (l1Index != nil) || (l2Index != nil) {
		minNode := getMinNode(l1Index, l2Index)

		// 第一个 Node —— 链表头
		if current == nil {
			result = minNode
			current = minNode
		} else {
			current.next = minNode
			current = current.next
		}

		if minNode == l1Index {
			l1Index = l1Index.next
		} else {
			l2Index = l2Index.next
		}
	}
	return
}

func getMinNode(node1 *Node, node2 *Node) *Node {
	if node1 == nil {
		return node2
	} else if node2 == nil {
		return node1
	} else {
		if node1.val <= node2.val {
			return node1
		} else {
			return node2
		}
	}
}
