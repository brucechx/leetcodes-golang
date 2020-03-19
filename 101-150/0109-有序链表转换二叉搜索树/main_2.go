package _109_有序链表转换二叉搜索树

func sortedListToBST2(head *ListNode) *TreeNode {
	nums := make([]int, 0)
	curr := head
	for curr != nil{
		nums = append(nums, curr.Val)
		curr = curr.Next
	}
	return sortedArrayToBST(nums, 0, len(nums) - 1)
}

/*
	有序数组转换二叉搜索树
 */
func sortedArrayToBST(nums []int, left, right int) *TreeNode{
	if left > right{
		return nil
	}
	mid := left + (right - left) >> 1
	root := &TreeNode{Val:nums[mid]}
	if left == right{
		return root
	}
	root.Left = sortedArrayToBST(nums, left, mid - 1)
	root.Right = sortedArrayToBST(nums, mid+1, right)
	return root
}
