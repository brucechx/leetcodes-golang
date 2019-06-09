package _51_数组中重复的数字

import (
	"sort"
)

// 排序
func Duplicate1(numbers []int) int{
	sort.Ints(numbers)
	for i:=1; i< len(numbers); i++{
		if numbers[i] == numbers[i-1]{
			return numbers[i]
		}
	}
	return -1
}

// map hash
func Duplicate2(numbers []int) int{
	numMap := make(map[int]bool)
	for _, v := range numbers{
		if _, ok:= numMap[v]; ok{
			return v
		}else {
			numMap[v] = true
		}
	}
	return -1
}


/**
利用元素数字都在0到n-1的范围内的特点，若不存在重复数字，则排序后的数字就在与其相同的索引值的位置，
即数字i在第i个位置。遍历的过程中寻找位置和元素不相同的值，并进行交换排序。时间复杂度O(n)，空间复杂度O(1)。
以数组｛2，3，1，0，2，5，3｝为例来分析找到重复数字的步骤。
数组的第0个数字（从0开始计数，和数组的下标保持一致）是2，与它的下标不想等，于是把它和下标为2的数字1交换，交换后的数组是｛1，3，2，0，2，5，3｝。
此时第0 个数字是1，仍然与它的下标不想等，继续把它和下标为1的数字3交换，得到数组｛0，1，2，3，2，5，3｝。
此时第0 个数字为0，接着扫描下一个数字，在接下来的几个数字中，下标为1，2，3的三个数字分别为1，2，3，他们的下标和数值都分别相等，因此不需要做任何操作。
接下来扫描下标为4的数字2.由于它的值与它的下标不登，再比较它和下标为2的数字。
注意到此时数组中下标为2的数字也是2，也就是数字2和下标为2和下标4的两个位置了，因此找到一个重复的数字。
*/
func Duplicate3(numbers []int) int {
	i := 0
	for i < len(numbers) {
		if i == numbers[i] {
			i++
		} else {
			if numbers[i] != numbers[numbers[i]] {
				numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
			} else {
				return numbers[i]
			}
		}
	}
	return -1
}