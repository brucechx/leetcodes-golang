package _014_最长公共前缀

func longestCommonPrefix(strs []string) string{
	short := shortest(strs)
	for i, r := range short{
		for j := range strs{
			if strs[j][i] != byte(r){
				return short[:i]
			}
		}
	}
	return short
}

func shortest(strs []string) string{
	if len(strs) == 0{
		return ""
	}

	res := strs[0]
	for _, s := range strs{
		if len(s) < len(res){
			res = s
		}
	}
	return res
}
