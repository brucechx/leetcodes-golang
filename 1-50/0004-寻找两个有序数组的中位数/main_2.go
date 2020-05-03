package _004_寻找两个有序数组的中位数

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	nums := combine(nums1, nums2)
	return medianOf(nums)
}

func combine(nums1, nums2 []int) []int{
	i, nums1Len := 0, len(nums1)
	j, nums2Len := 0, len(nums2)
	nums := make([]int, nums1Len+nums2Len)
	for k:=0; k<nums1Len+nums2Len; k++{
		if i == nums1Len || (i < nums1Len && j < nums2Len && nums1[i] > nums2[j]){
			nums[k] = nums2[j]
			j++
			continue
		}
		if j == nums2Len || (i < nums1Len && j < nums2Len && nums1[i] <= nums2[j]){
			nums[k] = nums1[i]
			i++
		}
	}
	return nums
}

func medianOf(nums []int) float64{
	numsLen := len(nums)
	if numsLen == 0 {
		return float64(0)
	}

	if numsLen % 2 == 0 {
		return float64(nums[numsLen/2] + nums[numsLen/2 - 1]) / 2.0
	}
	return float64(nums[numsLen/2])
}