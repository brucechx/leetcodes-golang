package _223_矩形面积

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	//如果有重叠矩形，该矩形 P,Q  X,Y 的计算方式如下：
	P := max(A, E)
	Q := max(B, F)
	X := min(C, G)
	Y := min(D, H)
	area := (C - A) * (D - B)   //先算一个，防止加法溢出
	//有重叠矩形判断条件，X,Y 在 P,Q 点的右上方
	if X > P && Y > Q {
		area -= (X - P) * (Y - Q)
	}
	return area + (G - E) * (H - F)
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}