package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	var res []*TreeNode
	if n <= 0 {
		return res
	}
	// 所有可能的中序遍历均一致
	inOrder := make([]int, n, n)
	for i:=1; i<=n; i++{
		inOrder[i-1] = i
	}
	// 基于中序遍历获取所有的前序遍历结果
	allPreOrder := getPreOrderByInOrder(inOrder)
	for _, preOrder := range allPreOrder{
		treeNode := rebuildTree(preOrder, inOrder)
		res = append(res, treeNode)
	}
	return res
}

func getPreOrderByInOrder(inOrder []int) [][]int{
	n := len(inOrder)
	if n <= 1{
		return [][]int{inOrder}
	}
	if n == 2{
		return [][]int{
			{inOrder[1], inOrder[0]},
			{inOrder[0], inOrder[1]},
		}
	}
	res := make([][]int, 0)
	for i:=0; i<n; i++{
		ls := getPreOrderByInOrder(inOrder[:i])
		rs := getPreOrderByInOrder(inOrder[i+1:])
		for _, l := range ls{
			for _, r := range rs{
				tmp := make([]int, 1, n)
				tmp[0] = inOrder[i]
				tmp = append(tmp, l...)
				tmp = append(tmp, r...)
				res = append(res, tmp)
			}
		}
	}
	return res
}


func rebuildTree(preOrder, inOrder []int) *TreeNode{
	if len(inOrder) == 0 || len(preOrder) == 0{
		return nil
	}
	root := &TreeNode{
		Val:   preOrder[0],
		Left:  nil,
		Right: nil,
	}
	rootIndex := findIndex(inOrder, preOrder[0])
	if rootIndex == -1{
		return nil
	}
	root.Left = rebuildTree(preOrder[1:rootIndex+1], inOrder[:rootIndex])
	root.Right = rebuildTree(preOrder[rootIndex+1:], inOrder[rootIndex+1:])
	return root
}

func findIndex(order []int, target int) int{
	for i, v := range order{
		if v == target{
			return i
		}
	}
	return -1
}
