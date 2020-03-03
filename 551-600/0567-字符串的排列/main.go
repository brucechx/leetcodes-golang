package _567_字符串的排列

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2){
		return false
	}
	s1Len := len(s1)
	left := 0
	s1Map := makeMap(s1)
	for right:=s1Len; right<=len(s2); right++{
		window := s2[left:right]
		if len(window) < s1Len{
			return false
		}
		// check window is valid
		if isValidWindow(window, s1Map){
			return true
		}
		left++
	}
	return false
}


func makeMap(s string) []int{
	res := make([]int, 26)
	for _, v := range s{
		res[v-'a']++
	}
	return res
}

func isValidWindow(window string, s1 []int) bool{
	winSlice := makeMap(window)
	if len(winSlice) != len(s1){
		return false
	}
	for i:=0; i<len(s1); i++{
		if winSlice[i] != s1[i]{
			return false
		}
	}
	return true
}
