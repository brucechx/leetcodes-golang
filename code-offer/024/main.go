package _24

// 二叉搜索树的后序遍历 左右根

func VerifySequenceBst(sequence []int) bool{

	if len(sequence) < 0{
		return false
	}

	root := sequence[len(sequence) - 1]

	i, j := 0, 0
	for ;i<len(sequence) - 1;i++{
		if sequence[i] > root{
			break
		}
	}
	j = i
	for ;j<len(sequence) - 1;j++{
		if sequence[j] < root{
			return false
		}
	}
	// 判断左子树
	left := true
	if i > 0{
		left = VerifySequenceBst(sequence[:i])
	}
	// 判断右子树
	right := true
	if j < len(sequence) - 1{
		right = VerifySequenceBst(sequence[i:len(sequence)-1])
	}
	return left && right
}
