package _031_下一个排列

/*
	1. 从右向左，应该是递增；找到a[i-1] < a[i]; 临界点为i-1;
	2. 逆转a[i:] 改为降序
	3. 再在a[i:] 中找到第一个不小于a[i-1]的数a[j]
	4. 进行交换
*/

func nextPermutation(nums []int)  {
	i := len(nums) - 1
	for i > 0 {
		// 得到临界点 i-1
		if nums[i-1] < nums[i]{
			break
		}
		i--
	}
	// 原先完全递增
	if i == 0 {
		reverseNum(nums, 0)
		return
	}else{
		reverseNum(nums, i)
	}
	// 从nums[i:] 找到第一个大于nums[i-1]的nums[j]
	j:=i
	for j<len(nums){
		if nums[j] > nums[i-1]{
			break
		}
		j++
	}
	// 交换nums[j]和nums[i-1]
	nums[i-1], nums[j] = nums[j], nums[i-1]
	return
}

func reverseNum(nums []int, i int) {
	// 逆转nums[i:]
	l:=i
	r:=len(nums)-1
	for l<r{
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
