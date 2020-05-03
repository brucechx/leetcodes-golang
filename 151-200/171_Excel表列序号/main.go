package _71_Excel表列序号

func titleToNumber(s string) int {
	res := 0

	for _, c := range s{
		v := c - 'A' + 1
		res = res * 26 + int(v)
	}
	return res
}

