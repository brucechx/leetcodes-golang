package _069_x的平方根

// 当 x≥2 时，它的整数平方根一定小于 x/2 且大于 0
func mySqrtBak(x int) int {
	if x < 2{
		return x
	}
	left := 2
	right := x / 2
	for left <= right{
		pivot := left + (right - left) / 2
		num := pivot * pivot
		switch  {
		case num > x:
			right = pivot - 1
		case num < x:
			left = pivot + 1
		default:
			return pivot
		}
	}
	return right
}

// 牛顿法
func mySqrt2(x int) int {
	if x < 2{
		return x
	}
	x0 := x
	x1 := (x0 + x/x0) / 2

	for (x0 - x1) < 1{
		x0 = x1
		x1 = (x0 + x/x0) / 2
	}
	return x1
}

func abs(x int) int{
	if x < 0 {
		return -x
	}else {
		return x
	}
}