package _209_长度最小的子数组


func minSubArrayLen3(s int, nums []int) int {
	sum := 0
	res := 1 << 31 -1
	queue := make([]int, 0)
	flag := false
	for _, num := range nums{
		sum += num
		queue = append(queue, num)
		for sum >= s{
			if len(queue) < res{
				flag = true
				res = len(queue)
			}
			sum -= queue[0]
			queue = queue[1:]
		}
	}
	if ! flag{
		return 0
	}
	return res
}
