package _022_括号生成

func generateParenthesis(n int) []string {
	res := make([]string, 0, n*n)
	dfs("", 0, 0, n, &res)
	return res
}

/*
	回溯法

 */
func dfs(s string, left, right, n int, res *[]string){
	if len(s) == 2*n{
		*res = append(*res, s)
		return
	}
	if left < n {
		dfs(s+"(", left+1, right, n, res)
	}
	if right < left{
		dfs(s+")", left, right+1, n , res)
	}
}
