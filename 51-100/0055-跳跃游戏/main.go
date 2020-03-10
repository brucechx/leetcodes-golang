package _055_跳跃游戏

func canJump(nums []int) bool {
	lastPosix := len(nums) - 1
	for i:=len(nums) -1; i>=0; i++{
		if nums[i] + i >= lastPosix{
			lastPosix = i
		}
	}
	return lastPosix == 0
}
