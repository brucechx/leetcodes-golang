package _215_数组中的第K个最大元素

/*
	选择排序
*/

func findKthLargest(nums []int, k int) int {
	for i:=0; i<k; i++{
		index := i
		for j:=i; j<len(nums); j++{
			if nums[j] > nums[index]{
				index = j
			}
		}
		if index != i{
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	return nums[k-1]
}
