package _078_子集

/*
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return

    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
 */

func subsets(nums []int) [][]int{
	res := make([][]int, 0)
	path := make([]int, 0)
	dfs(nums, &res, path, 0)
	return res
}

func dfs(nums []int, res *[][]int, path []int, pos int){
	tmp := make([]int, len(path))
	copy(tmp, path)
	*res = append(*res, tmp)
	for i:=pos; i<len(nums); i++{
		path = append(path, nums[i])
		dfs(nums, res, path, i+1)
		path = path[:len(path) - 1]
	}
}
