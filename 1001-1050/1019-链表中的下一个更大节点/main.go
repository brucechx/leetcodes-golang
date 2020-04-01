package _019_链表中的下一个更大节点

type ListNode struct {
	Val int
	Next *ListNode
}

/*
	单调递减栈实现
*/

func nextLargerNodes(head *ListNode) []int {
	stack := make([]int, 0) // 单调递减，保存索引下标
	data := make([]int, 0) // 索引对应的值
	result := make([]int, 0)
	index := 0 // 索引坐标
	for head != nil{
		result = append(result, 0)
		for len(stack) >0 && head.Val > data[stack[len(stack)-1]]{
			result[stack[len(stack) - 1]] = head.Val
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, index)
		data = append(data, head.Val)
		index++
		head = head.Next
	}
	return result[:index-1]
}
