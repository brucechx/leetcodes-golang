package _134_回文链表

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
	left, right := 0, len(nums) - 1
	for left < right{
		if nums[left] != nums[right]{
			return false
		}
		left++
		right--
	}
	return true
}
