package _62_序列化二叉树


//二叉树节点
type BinTreeNode struct {
	data   interface{}  //数据域
	parent *BinTreeNode //父节点
	lChild *BinTreeNode //左孩子
	rChild *BinTreeNode //右孩子
	height int          //以该结点为根的子树的高度
	size   int          //该结点子孙数(包括结点本身)
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


