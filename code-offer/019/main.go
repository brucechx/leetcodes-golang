package _19

type TreeNode struct {
	left *TreeNode
	right *TreeNode
	val int
}

func MirrorRecursively(root *TreeNode){
	if root == nil || (root.left == nil && root.right==nil){
		return
	}
	root.left, root.right = root.right, root.left
	MirrorRecursively(root.left)
	MirrorRecursively(root.right)
}
