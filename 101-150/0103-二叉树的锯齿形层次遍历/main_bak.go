package _103_二叉树的锯齿形层次遍历

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil{
		return res
	}
	currentLevelNodes := []*TreeNode{root}
	for len(currentLevelNodes) != 0 {
		var levelNodes []int
		var nextLevelNodes []*TreeNode
		for _, node := range currentLevelNodes{
			levelNodes = append(levelNodes, node.Val)
			if node.Left != nil{
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil{
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}
		res = append(res, levelNodes)
		currentLevelNodes = nextLevelNodes
	}
	// 偶数反转
	for i:=0; i<len(res); i++{
		if i % 2 == 0{
			continue
		}
		if len(res[i]) == 0 {
			continue
		}
		for j:=0; j< len(res[i])/2; j++{
			res[i][j], res[i][len(res[i]) - 1 -j] = res[i][len(res[i]) - 1 -j],res[i][j]
		}
	}
	return res
}

func zigzagLevelOrder2(root *TreeNode) [][]int {
	var res [][]int

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		// 出现了新的 level
		if level >= len(res) {
			res = append(res, []int{})
		}
		// 与 102 题相比，level 的奇偶不同，append 会不同。
		if level%2 == 0 {
			res[level] = append(res[level], root.Val)
		} else {
			temp := make([]int ,  len(res[level])+1)
			temp[0] = root.Val
			copy(temp[1:], res[level])
			res[level] = temp
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}
