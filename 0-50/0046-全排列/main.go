package _046_全排列

/*
for 选择 in 选择列表:
    # 做选择
    将该选择从选择列表移除
    路径.add(选择)
    backtrack(路径, 选择列表)
    # 撤销选择
    路径.remove(选择)
    将该选择再加入选择列表
*/

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	backtrack(nums, track, &res)
	return res
}

// 路径：记录在 track 中
// 选择列表：nums 中不存在于 track 的那些元素
// 结束条件：nums 中的元素全都在 track 中出现
func backtrack(nums []int, track []int, res *[][]int){
	// 结束条件
	if len(track) == len(nums){
		*res = append(*res, track)
	}
	for _, num := range nums{
		// 排除不合法的
		if exist(track, num){
			continue
		}
		// 做出选择
		track = append(track, num)
		// 进入下一层决策树
		backtrack(nums, track, res)
		// 取消选择
		track = track[:len(track)-1]
	}
}

func exist(track []int, num int) bool{
	for _, n := range track{
		if n == num{
			return true
		}
	}
	return false
}