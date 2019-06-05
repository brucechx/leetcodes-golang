package main

func Sum0(num int) int{
	return 0
}

func SumN(num int) int{
	fun := map[bool]func(num int) int{
		false: SumN,
		true: Sum0,
	}
	key := isZero(num)
	return num + fun[key](num - 1)
}

func isZero(num int) bool{
	return num == 0
}


