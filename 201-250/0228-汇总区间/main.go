package _228_汇总区间

import "fmt"

func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1{
		return []string{fmt.Sprintf("%d", nums[0])}
	}
	l := 0
	for r :=1 ;r < len(nums); r++{
		if nums[r] - nums[r-1] == 1{
			continue
		}
		if l == r {
			res = append(res, fmt.Sprintf("%d", nums[l]))
		}else {
			res = append(res, fmt.Sprintf("%d->%d", nums[l], nums[r-1]))
		}
		l = r
	}
	return res
}
