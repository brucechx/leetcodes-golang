package _58_二叉树的下一个结点

//二叉树
type binaryTree struct {
	root   *BinTreeNode //根节点
	height int
	size   int
}

func NewBinaryTree(root *BinTreeNode) *binaryTree {
	return &binaryTree{root: root}
}

//获得第一个与数据e相等的节点
func (b *binaryTree) Find(e interface{}) *BinTreeNode {
	if b.root == nil {
		return nil
	}
	p := b.root
	return isEqual(e, p)
}


func isEqual(e interface{}, node *BinTreeNode) *BinTreeNode {
	if e == node.GetData() {
		return node
	}

	if node.HasLChild() {
		lp := isEqual(e, node.GetLChild())
		if lp != nil {
			return lp
		}
	}

	if node.HasRChild() {
		rp := isEqual(e, node.GetRChild())
		if rp != nil {
			return rp
		}

	}

	return nil
}

func (b *binaryTree)InOrderNextNode(node *BinTreeNode) *BinTreeNode{
	if node == nil{
		return nil
	}
	var nextNode *BinTreeNode
	// 如果当前结点有右子树, 那么其中序遍历的下一个结点就是其右子树的最左结点
	if node.HasRChild(){
		//  找到右子树的最左孩子
		nextNode = node.GetRChild()
		for nextNode.HasLChild(){
			nextNode = nextNode.GetLChild()
		}
	}else if node.HasParent(){
		parentNode := node.GetParent()
		//  如果当前结点是其父结点的左子结点那么其中序遍历的下一个结点就是他的父亲结点
		if parentNode.GetLChild() == node{
			return parentNode
		}

		//  如果当前结点是其父结点的右子结点
		//  这种情况下其下一个结点应该是当前结点所在的左子树的根
		//  因此我们可以顺着其父节点一直向上遍历
		//  直到找到一个是它父结点的左子结点的结点
		curr := node
		for parentNode != nil && curr == parentNode.GetRChild(){
			curr = parentNode
			parentNode = parentNode.GetParent()
		}
		nextNode = parentNode
	}
	return nextNode
}
