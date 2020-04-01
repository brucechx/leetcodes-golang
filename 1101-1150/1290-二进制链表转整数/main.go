package _290_二进制链表转整数

type ListNode struct {
	Val int
	Next *ListNode
}

// 链表移到右侧一位，也即二进制左移一位
func getDecimalValue(head *ListNode) int {
	sum := 0
	for curr := head; curr != nil; curr = curr.Next{
		sum = sum << 1
		sum += curr.Val
	}
	return sum
}
