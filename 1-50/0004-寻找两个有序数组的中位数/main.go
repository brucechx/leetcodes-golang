package _004_寻找两个有序数组的中位数

// 二分查找法
/*
	i: nums1 分段
	j: nums2 分段
    左半部分 等于 右半部分 ==> i + j = m - i + n - j  ==> j = (m + n)/2 - i
		左半部分最大的值小于等于右半部分最小的值 max ( A [ i - 1 ] , B [ j - 1 ]）） <= min ( A [ i ] , B [ j ]））

	左半部分 比 右半部分 多一个 ==> i + j = m - i + n -j + 1 ==> j = (m + n + 1)/2 - i
		左半部分最大的值小于等于右半部分最小的值 max ( A [ i - 1 ] , B [ j - 1 ]）） <= min ( A [ i ] , B [ j ]））

	所以我们只需要保证 B [ j - 1 ] < = A [ i ] 和 A [ i - 1 ] <= B [ j ]

*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n{
		return findMedianSortedArrays(nums2, nums1)
	}
	// 保证m <= n
	iMin, iMax := 0, m
	for iMin <= iMax{
		i := iMin + (iMax - iMin) >> 1
		j := (m + n + 1) >> 1 - i
		if  j !=0 && i !=m && nums2[j-1] > nums1[i]{ // i 需要增大; i 增大，j 减小
			iMin = i + 1
		}else if i !=0 && j !=n && nums1[i-1] > nums2[j] { // i 需要减小
			iMax = i - 1
		}else { // 达到要求，将边界列出来单独考虑
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			}else if j == 0{
				maxLeft = nums1[i-1]
			}else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}
			if (m + n) & 1 == 1{ // 奇数，不需要考虑右半部分
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			}else if j == n {
				minRight = nums1[i]
			}else {
				minRight = min(nums2[j], nums1[i])
			}
			return (float64(maxLeft) + float64(minRight)) / 2.0 //如果是偶数的话返回结果
		}
	}
	return float64(0)
}

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}

func min(a, b int) int{
	if a < b{
		return a
	}
	return b
}