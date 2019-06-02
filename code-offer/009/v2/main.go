package v2

func Fib(n int) int{
	if n == 0 || n == 1{
		return n
	}
	tmp1, tmp2, fibN := 0, 1, 0
	for i:=2;i<=n;i++{
		fibN = tmp1 + tmp2
		tmp1 = tmp2
		tmp2 = fibN
	}
	return fibN
}
