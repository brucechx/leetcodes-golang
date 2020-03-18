package _222_完全二叉树的节点个数

import (
	"math"
)

/*
	利用完全二叉树特性:
	完全二叉树中，除了最后一层外，其余每层节点都是满的，并且最后一层的节点全部靠向左边。
*/

func countNodes3(root *TreeNode) int {
	if root == nil{
		return 0
	}
	d := treeDepth(root)
	if d == 0{
		return 1
	}
	// Last level nodes are enumerated from 0 to 2**d - 1 (left -> right).
	// Perform binary search to check how many nodes exist.
	left, right := 0, pow(2, d) - 1
	for left <= right{
		pivot := left + (right - left) // 2
		if exist(root, pivot, d){
			left = pivot + 1
		}else {
			right = pivot - 1
		}
	}
	// The tree contains 2**d - 1 nodes on the first (d - 1) levels
	// and left nodes on the last level.
	//fmt.Println(pow(2, d)-1, " ", left)
	return pow(2, d) - 1 + left
}

func pow2(a, b int) int{
	return int(math.Pow(float64(a), float64(b)))
}

func pow(a, b int) int{
	res := 1
	for b > 0 {
		res *= a
		b--
	}
	return res
}
// Return tree depth in O(d) time.
func treeDepth(root *TreeNode) int{
	d := 0
	for root.Left != nil{
		d++
		root = root.Left
	}
	return d
}

// Last level nodes are enumerated from 0 to 2**d - 1 (left -> right).
// Return True if last level node idx exists.
// Binary search with O(d) complexity.
func exist(node *TreeNode, idx, depth int) bool{
	left, right := 0, pow(2, depth) - 1
	for i:=0; i<depth; i++{
		pivot := left + (right - left) >> 1
		if idx <= pivot{
			node = node.Left
			right = pivot
		}else {
			node = node.Right
			left = pivot + 1
		}
	}
	return node != nil
}
