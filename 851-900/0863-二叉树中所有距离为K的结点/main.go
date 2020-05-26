package _863_二叉树中所有距离为K的结点

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 层序遍历
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if root == nil || target == nil{
		return []int{}
	}
	// 父子关系，可以向上层遍历
	parentMap := make(map[*TreeNode]*TreeNode)
	dfs(root, nil, parentMap)
	// 层级遍历队列
	queue := []*TreeNode{target}
	// 以遍历过的点
	seen := make(map[*TreeNode]bool)
	seen[target] = true
	// 开始层级遍历
	for len(queue) > 0 && K >0 {
		// 当前层
		size := len(queue)
		K--
		for size > 0{
			size--
			// 处理当前层
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil && ! seen[node.Left]{
				seen[node.Left] = true
				queue = append(queue, node.Left)
			}
			if node.Right != nil && ! seen[node.Right]{
				seen[node.Right] = true
				queue = append(queue, node.Right)
			}
			if parent, ok := parentMap[node]; ok && parent != nil && ! seen[parent]{
				seen[parent] = true
				queue = append(queue, parent)
			}
		}
	}
	result := make([]int, 0)
	for _, node := range queue{
		result = append(result, node.Val)
	}
	return result
}

func dfs(root, parent *TreeNode, parentMap map[*TreeNode]*TreeNode){
	if root == nil{
		return
	}
	parentMap[root] = parent
	dfs(root.Left, root, parentMap)
	dfs(root.Right, root, parentMap)
}
