package _043_字符串相乘

import (
	"fmt"
)

/*
//      0  1  2

		1  2  3
           4  5
    -------------
           1  5  index(i+j, i+j+1)
        1  0
     0  5
    --------------
        1  2
     0  8
  0  4

//0  1  2  3  4
 ====================
 */

func multiply(num1 string, num2 string) string {
	if num2 == "0" || num1 == "0"{
		return "0"
	}
	res := make([]int, len(num1)+len(num2))

	for i:=len(num2)-1; i>=0; i--{
		for j:=len(num1)-1; j>=0; j--{
			n2 := int(num2[i] - '0')
			n1 := int(num1[j] - '0')
			sum := res[i+j+1] + n1 * n2
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	var resStr string
	for i, v := range res{
		if i==0 && v == 0{
			continue
		}
		resStr += fmt.Sprintf("%d", v)
	}
	return resStr
}
