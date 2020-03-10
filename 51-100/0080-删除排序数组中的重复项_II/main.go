package _080_删除排序数组中的重复项_II

func removeDuplicates(nums []int) int {
	i := 2
	for j:=i; j<len(nums); j++{
		if nums[i-2] != nums[j]{
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
