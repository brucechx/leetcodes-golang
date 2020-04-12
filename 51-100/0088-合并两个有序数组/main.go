package _088_合并两个有序数组

func merge(nums1 []int, m int, nums2 []int, n int)  {
	// 双指针法，从后向前
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1 // new nums1

	for p1 >=0 && p2 >=0{
		if nums1[p1] > nums2[p2]{
			nums1[p] = nums1[p1]
			p1--
			p--
		}else {
			nums1[p] = nums2[p2]
			p2--
			p--
		}
	}

	// 拷贝nums2 剩下的
	copy(nums1[:p2+1], nums2[:p2+1])
}
