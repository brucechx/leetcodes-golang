package _171_从链表中删去总和值为零的连续节点

type ListNode struct {
	Val int
	Next *ListNode
}

/*
	前缀和
*/

func removeZeroSumSublists(head *ListNode) *ListNode {
	sumMap := make(map[int]*ListNode)
	newHead := head
	sum := 0
	for p:=head; p!=nil; p=p.Next{
		sum += p.Val
		// 情况一
		if sum == 0 {
			sumMap = make(map[int]*ListNode)
			newHead = p.Next
			continue
		}
		// 情况二
		if v, ok := sumMap[sum]; !ok || v == nil{
			sumMap[sum] = p
			continue
		}
		// 情况三
		lastNode := sumMap[sum]
		tmpSum := sum
		for c:=lastNode.Next; c != p; c = c.Next{
			tmpSum += c.Val
			sumMap[tmpSum] = nil
		}
		// 删除
		lastNode.Next = p.Next
	}
	return newHead
}
