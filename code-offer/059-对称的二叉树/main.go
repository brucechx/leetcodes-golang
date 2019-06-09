package _59_对称的二叉树

//二叉树节点
type BinTreeNode struct {
	data   interface{}  //数据域
	parent *BinTreeNode //父节点
	lChild *BinTreeNode //左孩子
	rChild *BinTreeNode //右孩子
	height int          //以该结点为根的子树的高度
	size   int          //该结点子孙数(包括结点本身)
}

func NewBinTreeNode(e interface{}) *BinTreeNode {
	return &BinTreeNode{data: e, size: 1}
}


//二叉树
type binaryTree struct {
	root   *BinTreeNode //根节点
	height int
	size   int
}

func NewBinaryTree(root *BinTreeNode) *binaryTree {
	return &binaryTree{root: root}
}

func (b *binaryTree) isSymmetrical() bool{
	return isSymmetrical(b.root, b.root)
}

func isSymmetrical(leftNode, rightNode *BinTreeNode) bool{
	if leftNode == nil && rightNode == nil{
		return true
	}
	if leftNode == nil || rightNode == nil{
		return false
	}
	if leftNode.data != rightNode.data {
		return false
	}
	return isSymmetrical(leftNode.lChild, rightNode.rChild) &&
		isSymmetrical(leftNode.rChild, rightNode.lChild)
}