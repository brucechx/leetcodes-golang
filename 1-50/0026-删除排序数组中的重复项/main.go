package _026_删除排序数组中的重复项

func removeDuplicates(a []int) int{
	if len(a) == 0 {
		return 0
	}
	i := 0 // 慢指针
	for j:=1; j< len(a); j++{
		if a[i] == a[j]{
			continue
		}
		i ++
		a[i] = a[j]
	}
	return i + 1
}
