package _041_缺失的第一个正数

/*
+ 检查 1 是否存在于数组中。如果没有，则已经完成，1 即为答案。
+ 如果 nums = [1]，答案即为 2 。
+ 将负数，零，和大于 n 的数替换为 1 。
+ 遍历数组。当读到数字 a 时，替换第 a 个元素的符号。
+ 注意重复元素：只能改变一次符号。由于没有下标 n ，使用下标 0 的元素保存是否存在数字 n。
+ 再次遍历数组。返回第一个正数元素的下标。
+ 如果 nums[0] > 0，则返回 n 。
+ 如果之前的步骤中没有发现 nums 中有正数元素，则返回n + 1。
 */

func firstMissingPositive(nums []int) int {
	n := len(nums)
	contains := 0
	for _, v := range nums{
		if v == 1{
			contains++
			break
		}
	}
	// 1. 1 不存在
	if contains == 0{
		return 1
	}
	// 2. nums = [1]
	if n == 1{
		return 2
	}
	// 3. 负数，零，和大于 n 的数替换为 1
	for i:=0; i<n; i++{
		if nums[i] <= 0 || nums[i] > n{
			nums[i] = 1
		}
	}
	// 4. 遍历数组。当读到数字 a 时，替换第 a 个元素的符号。
	for i:=0; i< n; i++{
		a := abs(nums[i])
		if a == n{
			nums[0] = - abs(nums[0])
		}else {
			nums[a] = - abs(nums[a])
		}
	}
	// 5. 现在第一个正数的下标, 就是第一个缺失的数
	for i:=1; i<n; i++{
		if nums[i] > 0{
			return i
		}
	}
	// 6.
	if nums[0] > 0{
		return n
	}
	// 7.
	return n + 1
}

func abs(i int) int{
	if i < 0{
		return -i
	}
	return i
}