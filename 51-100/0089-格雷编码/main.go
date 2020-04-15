package _089_格雷编码

/*
所以如果知道了 n = 2 的解的话，如果是 { 0, 1, 3, 2}，
那么 n = 3 的解就是 { 0, 1, 3, 2, 2 + 4, 3 + 4, 1 + 4, 0 + 4 }，即 { 0 1 3 2 6 7 5 4 }。
之前的解直接照搬过来，然后倒序把每个数加上 1 << ( n - 1) 添加到结果中即可。
*/

func grayCode(n int) []int {
	gray := make([]int, 1) // n=0 ==> [0]
	for i:=0; i<n; i++{
		add := 1 << i
		for j:=len(gray)-1; j>=0; j--{
			gray = append(gray, gray[j]+add)
		}
	}
	return gray
}
