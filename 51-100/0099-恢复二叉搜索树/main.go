package main
/*
	 所谓简单的O(n)个空间的思路：
	 中序遍历得到数组，比如[1,3,2,4],在用一个hashmap<int,TreeNode>保存对象.
	 然后对数组排序，得到[1,2,3,4],对比得到2和3两个节点不一样，交换一下值即可。
	 下面实现进阶的常数空间思想：
	 在纸上写几个数组观察一下，可以发现，只需要找到三个结点（prev_first,first,second），就能进行交换，
     且只交换值是最方便的。中序遍历，寻找第一个递减的节点，用prev记录前一个节点。
	 prev_first记录frist前驱，first记录现在。
	 再次寻找第二个递减的节点，second记录。若是找到了second，交换prev_first和second，
     若是没找到，交换prev_first和first

	// 如果要调整 [1, 2, 6, 4, 5, 3, 7] 中错位的 6 和 3， 6 为first 3 为second
	// 其实是把 [6, 4] 中的较大值与 [5, 3] 中的较小值交换。这时，有两组错序。
	// 但是，还有可能是
	// [1, 3, 2] 中错位的 3 和 2，只有一组错序的 [3, 2]
	// 以下的两个 if 语句，可以兼容以上两种情况

     使用中序遍历，交换下降的值
*/

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var firstNode, secondNode, prev *TreeNode
	// 中序遍历
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left != nil {
			dfs(node.Left)
		}
		// 有下降
		if prev != nil && prev.Val > node.Val {
			if firstNode == nil {
				firstNode = prev
			}

			if firstNode != nil {
				secondNode = node
			}
		}
		prev = node
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	// 中序遍历
	dfs(root)
	firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
}




