package _109_有序链表转换二叉搜索树

type ListNode struct {
    Val int
    Next *ListNode
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	mid, pre := findMidNode(head)
	if mid == nil{
		return nil
	}
	root := &TreeNode{Val:mid.Val}
	if mid == head{
		return root
	}
	pre.Next = nil
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(mid.Next)
	return root
}

func findMidNode(head *ListNode) (mid *ListNode, midPre *ListNode){
	if head == nil || head.Next == nil{
		return head, nil
	}
	pre := &ListNode{Next:head}
	slow, fast := head, head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		pre  = pre.Next
		fast = fast.Next.Next
	}
	return slow, pre
}

