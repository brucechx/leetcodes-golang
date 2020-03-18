package _255_验证前序遍历序列二叉搜索树

/*
	二叉搜索树的特性
	1. 左子树递减， 右子树递增
	2. 局部递减，总体递增
	//
	如果递减趋势结束后的下一个值，比递增趋势的第一个值小，则报错
*/

func verifyPreorder(preorder []int) bool {
	stack := make([]int, 0)
	minNum := -1 << 31
	for _, val := range preorder{
		// 如果递减趋势结束后的下一个值，比递增趋势的第一个值小，则报错
		if val < minNum{
			return false
		}
		for len(stack) > 0 && val > stack[len(stack) - 1]{
			// 递增趋势的第一个值
			minNum = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, val)
	}
	return true
}
