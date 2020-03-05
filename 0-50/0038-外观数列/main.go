package _038_外观数列


import "fmt"

var res = []string{"1", "1"}
func countAndSay(n int) string {
	if n < len(res) {
		return res[n]
	}

	// init result
	for i := len(res); i <= n; i++ {
		res = append(res, countInt(i))
	}
	fmt.Println("res=", res)
	return res[n]
}

func countInt(n int) string {
	countStr := ""
	nStr := res[n-1]
	for i := 0; i < len(nStr); i++ {
		cnt := 1
		for i+1 < len(nStr) && nStr[i+1] == nStr[i] {
			cnt++
			i++
		}
		countStr += fmt.Sprint(cnt) + string(nStr[i])
	}

	return countStr
}
