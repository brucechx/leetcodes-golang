package _654_最大二叉树

// 单调栈法
func constructMaximumBinaryTree2(nums []int) *TreeNode {
	// stack[i+1:] 中所有的节点都是 stack[i] 右边的节点
	// 即，stack 中的 Val 值应该始终是递减的
	// 在遍历 nums 后，stack[i].Right == stack[i+1]
	stack := make([]*TreeNode, 1, len(nums))
	// 先放这个节点到 stack 可以简化后面的逻辑结构
	// nums 生成的 Tree 肯定是这个节点的右节点
	stack[0] = &TreeNode{Val: 1<<63 - 1}
	for _, num := range nums{
		node := &TreeNode{Val:num}
		// 栈顶元素大于num
		if stack[len(stack)-1].Val > num{
			stack = append(stack, node)
			continue
		}
		// 栈顶元素小于num；缩小单调栈，找到第一个比num 大的作为栈顶
		for len(stack)-2 >=0 && stack[len(stack) - 2].Val < num{
			// 此时可以确定，stack中，上一项是下一项的 Right 节点
			stack[len(stack)-2].Right = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		// 此时，stack 中最后一项是 node 的 Left 节点
		node.Left = stack[len(stack)-1]

		// 把 node 放入 stack
		stack[len(stack)-1] = node
	}
	// 连接父子关系
	for i:=len(stack)-1; i>=1; i--{
		stack[i-1].Right = stack[i]
	}
	return stack[1]
}
