package _191_位1的个数

func hammingWeight(num uint32) int {
	sum := 0
	for num != 0{
		sum ++
		// 不段将最后一个1消0
		num &= num - 1
	}
	return sum
}

