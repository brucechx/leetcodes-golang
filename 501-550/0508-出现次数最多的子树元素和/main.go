package _508_出现次数最多的子树元素和

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
	1. 递归求和，map 保存
	2. 遍历map 求最大和
*/

func findFrequentTreeSum(root *TreeNode) []int {
	allSum := make(map[int]int)
	dfs(root, allSum)
	// get max sum
	max := -1
	res := make([]int, 0)
	// 针对所有子树和，求出最大集
	for key, val := range allSum{
		if max > val{
			continue
		}
		if max < val{
			max = val
			res = res[0:0]
		}
		res = append(res, key)
	}
	return res
}

// 递归求出所有子树和
func dfs(root *TreeNode, allSum map[int]int) int{
	if root == nil{
		return 0
	}
	s := root.Val + dfs(root.Left, allSum) + dfs(root.Right, allSum)
	allSum[s]++
	return s
}
