package _169_多数元素

func majorityElement(nums []int) int {
	curr, sum := 0, 0
	for _, num := range nums{
		if sum == 0{
			curr = num
		}
		if curr == num{
			sum++
		}else {
			sum--
		}
	}
	return curr
}
