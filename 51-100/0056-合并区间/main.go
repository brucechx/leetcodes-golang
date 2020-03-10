package _056_合并区间

import "sort"

type pair [][]int

func (p pair) Len() int{
	return len(p)
}

func (p pair) Swap(i, j int){
	p[i], p[j] = p[j], p[i]
}

func (p pair) Less(i, j int) bool{
	return p[i][0] < p[j][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1{
		return intervals
	}
	sort.Sort(pair(intervals))
	res := make([][]int, 0)
	for i := range intervals{
		n := len(res)
		if n == 0 || intervals[i][0] > res[n-1][1]{
			res = append(res, intervals[i])
		}else {
			res[n-1][1] = max(res[n-1][1], intervals[i][1])
		}
	}
	return res
}

