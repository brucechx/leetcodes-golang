package _198_打家劫舍

// 动态规划
/*
f(k):前k个房屋的最大值； Ai: 第i个房屋的最大值

f(k) = max(f(k-2)+Ak, f(k-1))
- 抢第k个Ak;
- 不抢第k个
*/

func rob(nums []int) int {
	preMax := 0 // f(k-2) + Ak
	currMax := 0 // f(k-1)
	for _, num := range nums{
		tmp := currMax
		currMax = max(preMax+num, currMax)
		preMax = tmp
	}
	return currMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
