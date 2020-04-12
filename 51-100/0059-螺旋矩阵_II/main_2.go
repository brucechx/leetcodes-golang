package _059_螺旋矩阵_II

func generateMatrix2(n int) [][]int {
	if n == 0 {
		return nil
	}

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	max := n * n
	next := nextFunc(n)

	for i := 1; i <= max; i++ {
		x, y := next()
		res[x][y] = i
	}

	return res
}

func nextFunc(n int) func() (int, int) {
	top, down := 0, n-1
	left, right := 0, n-1
	x, y := 0, -1
	dx, dy := 0, 1
	return func() (int, int) {
		x += dx
		y += dy
		switch {
		case y+dy > right: // top --> bottom x++ y不变
			top++
			dx, dy = 1, 0
		case x+dx > down: // right --> left  x不变 y--
			right--
			dx, dy = 0, -1
		case y+dy < left: // bottom --> top x-- y不变
			down--
			dx, dy = -1, 0
		case x+dx < top: // left --> right x 不变 y++
			left++
			dx, dy = 0, 1
		}
		return x, y
	}
}
