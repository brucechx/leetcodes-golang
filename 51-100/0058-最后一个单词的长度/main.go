package _058_最后一个单词的长度

func lengthOfLastWord(s string) int {
	if len(s) == 0{
		return 0
	}
	res := 0
	for i:= len(s)-1; i>=0; i--{
		if s[i] == ' '{
			if res != 0 {
				return res
			}
			continue
		}
		res++
	}
	return res
}
