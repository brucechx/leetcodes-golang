package main

import (
	"github.com/pkg/errors"
	"fmt"
)

func main() {
	nums := []int{3, 4, 5, 9}
	target := 7
	res := twoSum(nums, target)

	fmt.Println("res: ", res)

}


func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	for i, v := range nums{
		t := target - v
		if _, ok := tmp[t]; ok {
			return []int{tmp[t], i}
		}
		tmp[v] = i
	}
	panic(errors.New("no res"))
}
