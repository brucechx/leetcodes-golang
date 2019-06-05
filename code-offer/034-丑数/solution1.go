package _34_丑数

// 暴力破解法
func IsUglyNumber(num int) bool{
	for num % 2 == 0{
		num /= 2
	}

	for num % 3 == 0{
		num /= 3
	}

	for num % 5 == 0{
		num /= 5
	}

	return num == 1
}

func GetUglyNumberSolution1(index int) int{
	count := 0
	number := 1

	for {
		if IsUglyNumber(number){
			count ++
		}

		if count == index{
			return number
		}
		number++
	}
	return number
}