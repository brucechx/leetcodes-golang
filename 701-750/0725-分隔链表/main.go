package _725_分隔链表

type ListNode struct {
	Val int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	size := listLength(root)
	remainder := size % k // 均分后余量
	res := make([]*ListNode, k)
	i := 0
	for {
		res[i] = root
		i++
		if i == k{
			break
		}
		l := size / k // 每个桶大小
		if remainder > 0{ // 如果均分有余量，前面桶多一个
			l++
			remainder --
		}
		for l > 1 && root != nil{
			root = root.Next
			l --
		}
		if root == nil{ // size < k
			break
		}
		// 断开连接
		root.Next, root = nil, root.Next
	}
	return res
}

func listLength(root *ListNode) int{
	length := 0
	for root != nil{
		length++
		root = root.Next
	}
	return length
}