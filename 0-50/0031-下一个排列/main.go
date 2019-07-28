package _031_下一个排列

/*
	1. 从后往前找到最长的降序排列
    2. 把此降序排列转成升序排列
    3. 把序列前的元素，与序列中，第一个大于他的元素互换。
 */
/*
 	从后往前遍历，找到第一个下降的值nums[i]，
 	和之后刚刚大于这个值的nums[j]交换位置。
 	再对i之后的进行排序。
*/


// 逆转 a[l:]
func reverse(a []int, l int) {
	r := len(a) - 1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}

// 返回 a[l:] 中 > target 的最小值的索引号
// a[l:] 是一个 递增 数列
func search(a []int, l, target int) int {
	r := len(a) - 1
	l--
	for l+1 < r {
		mid := (l + r) / 2
		if target < a[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func nextPermutation(a []int) [] int{
	left := len(a) - 2
	for 0 <= left && a[left] >= a[left+1] {
		left--
	}

	// 此时 a[left+1:] 是一个 递减 数列

	reverse(a, left+1)

	if left == -1 {
		return a
	}

	// 此时 a[left+1:] 是一个 递增 数列

	right := search(a, left+1, a[left])
	a[left], a[right] = a[right], a[left]
	return a
}
