package _025_K个一组翻转链表

/*
	不足k个也反转
*/

func reverseKGroup2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	kNode := getKNode(head, k)
	if kNode != nil{
		tmp := kNode.Next
		kNode.Next = nil
		head, kNode = reverseNode(head)
		kNode.Next = reverseKGroup2(tmp, k)
	}else {
		head, _ = reverseNode(head)
	}
	return head
}

func getKNode(head *ListNode, k int) *ListNode{
	for k > 1 && head != nil{
		head = head.Next
		k--
	}
	return head
}

func reverseNode(head *ListNode) (newHeadNode, newTailNode *ListNode){
	var pre, curr *ListNode
	curr = head
	for curr != nil{
		pre, curr, curr.Next = curr, curr.Next, pre
	}
	return pre, head
}