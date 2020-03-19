package _234_回文链表

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	nums := make([]int, 0)
	for head != nil{
		nums = append(nums, head.Val)
		head = head.Next
	}
	l, r := 0, len(nums)-1
	for l < r{
		if nums[l] != nums[r]{
			return false
		}
		l++
		r--
	}
	return true
}
