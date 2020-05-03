package _078_子集

func subsets2(nums []int) [][]int {
	res := make([][]int, 1, 1024)
	for _, n := range nums{
		for _, r := range res{
			res = append(res, append([]int{n}, r...))
		}
	}
	return res
}
