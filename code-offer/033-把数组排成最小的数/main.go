package _33_把数组排成最小的数

import (
	"strings"
	"fmt"
	"sort"
)

type IntArray []int

func (a IntArray) Len() int{
	return len(a)
}

func (a IntArray) Less(i, j int) bool{
	return strings.Compare(fmt.Sprintf("%d%d", a[i], a[j]), fmt.Sprintf("%d%d", a[j], a[i])) < 0
}

func (a IntArray) Swap(i, j int){
	a[i], a[j] = a[j], a[i]
}


func PrintMinNums(numbers []int) string{
	if len(numbers) == 0{
		return ""
	}
	array := IntArray(numbers)
	sort.Sort(array)
	s := ""
	for _, v := range array{
		s = fmt.Sprintf("%s%d", s, v)
	}
	return s
}
