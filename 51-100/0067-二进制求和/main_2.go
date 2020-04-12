package _067_二进制求和


func addBinary2(a, b string) string{
	const  base  = 2
	m, n := len(a), len(b)
	if m > n {
		return addBinary2(b, a)
	}
	buff := make([]byte, n + 1)
	carry := 0
	for i, j := n-1, m-1; i>=0; i--{
		if j >= 0 {
			carry += int(a[j] - '0')
			j--
		}
		carry += int(b[i] - '0')
		buff[i+1] = byte(carry%base + '0')
		carry /= base
	}
	if carry == 0{
		return string(buff[1:])
	}
	buff[0] = '1'
	return string(buff)
}
