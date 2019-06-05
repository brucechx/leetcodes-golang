package _44_扑克牌顺子

func IsContinuous(array []int) bool{
	if len(array) != 5{
		return false
	}

	min, max := 13, 0
	flagNum := make([]int, 13)
	for _, v := range array{
		//  牌只能在0~13之间
		if v < 0 || v > 13{
			return false
		}
		//  0用来答题任何牌，因此不能参与最大最小牌的比对
		if v == 0{
			continue
		}

		//  如果flag的第num位非0, 说明num重复
		if flagNum[v] != 0{
			return false
		}else {
			flagNum[v] = 1
		}

		//  寻找最大最小的牌
		if v > max{
			max = v
		}
		if v < min{
			min = v
		}

		//  如果最大值和最小值的差值大于4, 那么必应不能补齐
		if max - min > 4{
			return false
		}
	}
	return true
}
