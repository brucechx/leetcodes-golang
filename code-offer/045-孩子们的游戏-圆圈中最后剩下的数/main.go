package _45_孩子们的游戏_圆圈中最后剩下的数

func LastRemainingSolution(n, m int)  int{
	if n < 1 || m < 1{
		return -1
	}
	last := 0
	for i:=2; i <=n; i++{
		last = (last + m)%i
	}
	return last
}

func LastRemainingSolution2(n, m int) int{
	if n < 1 || m < 1 {
		return -1
	}else if n == 1{
		return 0
	}else {
		// F[n] = (F[n - 1] + m) % n
		return (LastRemainingSolution2(n-1,m)+m)%n
	}
}
