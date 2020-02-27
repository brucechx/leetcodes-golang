package _113_路径总和_II

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 方法一
func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	var path []int

	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, level, sum int) {
		if node == nil {
			return
		}

		// 根据 level 来更新 path
		if level >= len(path) {
			path = append(path, node.Val)
		} else {
			path[level] = node.Val
		}
		sum -= node.Val

		// 到达 leaf
		// 并且，此条路径符合要求
		if node.Left == nil && node.Right == nil && sum == 0 {
			temp := make([]int, level+1)
			// copy 不会复制 path[len(temp):]，如果 len(path) > len(temp) 的话
			copy(temp, path)
			res = append(res, temp)
		}

		dfs(node.Left, level+1, sum)
		dfs(node.Right, level+1, sum)
	}

	dfs(root, 0, sum)

	return res
}


// 方法二
func pathSum2(root *TreeNode, sum int) [][]int {
	var res [][]int
	traverse(root, 0, []int{}, sum, &res)
	return res
}


func traverse(node *TreeNode, parentSum int, parentPath []int, target int, result *[][]int) {
	if node == nil {
		return
	}

	currentPath := append([]int{}, parentPath...)
	currentPath = append(currentPath, node.Val)
	currentSum := parentSum + node.Val

	// is leaf
	if node.Left == nil && node.Right == nil && currentSum == target {
		*result = append(*result, currentPath)
		return
	}

	if node.Left != nil {
		traverse(node.Left, currentSum, currentPath, target, result)
	}

	if node.Right != nil {
		traverse(node.Right, currentSum, currentPath, target, result)
	}
}