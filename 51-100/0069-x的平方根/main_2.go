package _069_x的平方根

/*
	求浮点数x的平方根y，精度误差e。
*/

func mySqrtFloat(x, e float32) float32{
	if x < 2{
		return x
	}
	left, right := float32(0), x/2
	var mid float32
	for right - left > e || right - left < -e {
		mid = left + (right - left) / 2
		num := mid * mid
		if num - x  > e {
			right = mid
		}else if num - x < -e{
			left = mid
		}else {
			return mid
		}
	}
	return mid
}

