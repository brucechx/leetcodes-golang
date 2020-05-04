package _654_最大二叉树

type TreeNode struct {
 	Val int
  	Left *TreeNode
  	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0{
		return nil
	}
	rootIndex := getMax(nums)
	root := &TreeNode{Val:nums[rootIndex]}
	root.Left = constructMaximumBinaryTree(nums[:rootIndex])
	root.Right = constructMaximumBinaryTree(nums[rootIndex+1:])
	return root
}

func getMax(nums []int) int{
	maxIndex := 0
	for i:=0; i<len(nums); i++{
		if nums[i] > nums[maxIndex]{
			maxIndex = i
		}
	}
	return maxIndex
}
