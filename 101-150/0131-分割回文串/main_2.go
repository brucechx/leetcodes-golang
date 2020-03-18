package _131_分割回文串

func partition2(s string) [][]string{
	res := make([][]string, 0)
	path := make([]string, 0)
	dfs2(s, &res, path, 0)
	return res
}

func dfs2(s string, res *[][]string, path []string, start int){
	if start == len(s){
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i:=start; i<len(s); i++{
		if ! par(s[start:i+1]){
			continue
		}
		path = append(path, s[start:i+1])
		dfs2(s, res, path, i+1)
		path = path[:len(path) - 1]
	}
}
