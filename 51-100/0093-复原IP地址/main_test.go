package _093_复原IP地址

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T){
	//input := "25525511135"
	input := "1111"
	res := restoreIpAddresses(input)
	fmt.Println(res)
}
