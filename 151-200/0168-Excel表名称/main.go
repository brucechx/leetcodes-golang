package _168_Excel表名称

func convertToTitle(n int) string {
	var ans []byte
	for ;n>0;n = n/26{
		n--
		ans = append([]byte{'A' + byte(n%26)},ans...)
	}
	return string(ans)
}
