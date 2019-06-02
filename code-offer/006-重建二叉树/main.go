package _06_重建二叉树

/*
	前序遍历： 根左右
	中序遍历： 左根右
    后序遍历： 左右根
 */

type Node struct {
	left  *Node
	right *Node
	value int
}

func (node *Node) PreOrderTraversal(result []int) []int {
	if node != nil {
		result = append(result, node.value)
	}
	if node.left != nil {
		result = node.left.PreOrderTraversal(result)
	}
	if node.right != nil {
		result = node.right.PreOrderTraversal(result)
	}

	return result
}


func (node *Node) InOrderTraversal(result []int) []int {
	if node.left != nil {
		result = node.left.PreOrderTraversal(result)
	}
	if node != nil {
		result = append(result, node.value)
	}
	if node.right != nil {
		result = node.right.PreOrderTraversal(result)
	}

	return result

}
func (node *Node) BreadthFirstTraversal(result []int) []int {
	queue := []*Node{node}
	var n *Node

	for len(queue) > 0 {
		n = queue[0]
		queue = queue[1:]

		result = append(result, n.value)

		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}

	return result
}


func RebuildBinaryTree(preOrderTraversalResult, inOrderTraversalResult []int) *Node{
	if len(preOrderTraversalResult) == 0{
		return nil
	}
	return rebuild(preOrderTraversalResult, inOrderTraversalResult,
		0, len(preOrderTraversalResult) - 1,
			0, len(inOrderTraversalResult) - 1)
}

func sliceIndex(nums []int, target int) int{
	for i, v := range nums{
		if v == target{
			return i
		}
	}
	return -1
}
func rebuild(preOrderTraversalResult, inOrderTraversalResult []int,
		preOrderStartIndex, preOrderEndIndex,
		inOrderStartIndex, inOrderEndIndex int) (rootNode *Node){
	// 基于前序遍历，找到根节点
	rootNodeVal := preOrderTraversalResult[preOrderStartIndex]
	rootNode = &Node{
		value:rootNodeVal,
		left:nil,
		right:nil,
	}

	// 递归终止条件：只有一个节点
	if preOrderStartIndex == preOrderEndIndex {
		return
	}

	//  从中序遍历中找到根节点位置
	rootIndexInOrder := sliceIndex(inOrderTraversalResult, rootNodeVal)
	if rootIndexInOrder == -1 {
		return
	}

	leftTreeNodeNum := rootIndexInOrder - inOrderStartIndex
	leftPreOrderEndIndex := preOrderStartIndex + leftTreeNodeNum
	// 因为前序遍历结果是 [根节点, 左子树结果, 右子树结果]，所以
	// preOrderTraversalResult[preOrderStartIndex+1: leftPreOrderEndIndex] 是左子树的前序遍历结果
	// 因为中序遍历结果是 [左子树结果, 根节点, 右子树结果]，所以
	// inOrderTraversalResult[inOrderStartIndex: rootIndexOfInOrderTraversalResult-1] 是左子树的中序遍历结果
	// 下面右子树的同理
	rootNode.left = rebuild(
		preOrderTraversalResult,
		inOrderTraversalResult,
		preOrderStartIndex+1, // 排除当前 root 节点
		leftPreOrderEndIndex,
		inOrderStartIndex,
		rootIndexInOrder-1,
	)
	nonRootNodeNum := preOrderEndIndex - preOrderStartIndex
	rightTreeNodeNum := nonRootNodeNum - leftTreeNodeNum
	if rightTreeNodeNum > 0 {
		rootNode.right = rebuild(
			preOrderTraversalResult,
			inOrderTraversalResult,
			leftPreOrderEndIndex+1,
			preOrderEndIndex,
			rootIndexInOrder+1,
			inOrderEndIndex,
		)
	}

	return
}

