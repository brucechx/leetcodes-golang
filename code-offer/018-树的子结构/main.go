package _18_树的子结构

type Node struct {
	left *Node
	right *Node
	val int
}

func IdenticalTrees(tree1, tree2 *Node) bool{
	result := false
	if tree1 != nil && tree2 != nil{
		if tree1.val == tree2.val{
			result = doTree1HasTree2(tree1, tree2)
		}
		if ! result{
			result = IdenticalTrees(tree1.left, tree2)
		}
		if ! result{
			result = IdenticalTrees(tree1.right, tree2)
		}
	}
	return result
}

func doTree1HasTree2(tree1, tree2 *Node) bool{
	if tree2 == nil{
		return true
	}

	if tree1 == nil{
		return false
	}

	if tree1.val != tree2.val{
		return false
	}
	return doTree1HasTree2(tree1.left, tree2.left) && doTree1HasTree2(tree1.right, tree2.right)
}

