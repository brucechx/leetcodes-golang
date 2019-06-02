package _22_栈的压入弹出序列

func IsPopOrder2(pushV []int, popV []int) bool {
	if len(pushV) == 0 || len(popV) == 0 || len(pushV) != len(popV) {
		return false
	}
	var s Stack
	currentPop := 0
	for _, v := range pushV {
		s.Push(v)
		for !s.isEmpty() {
			top := s.Top()
			if top.(int) == popV[currentPop] {
				s.Pop()
				currentPop++
			} else {
				break
			}
		}
	}
	if currentPop == len(popV) {
		return true
	} else {
		return false
	}
}