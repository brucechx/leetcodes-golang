package _017_电话号码的字母组合

var m = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if digits == "" {
		return res
	}
	dfs(&res, 0, "", digits)
	return res
}

// 类似全排列
func dfs(res *[]string, index int, path, digits string){
	if index == len(digits){
		*res = append(*res, path)
		return
	}
	letters := m[digits[index]]
	for _, l := range letters{
		dfs(res, index+1, path+l, digits)
	}
}
