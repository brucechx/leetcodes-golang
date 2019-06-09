package _60_把二叉树打印成多行

import "container/list"

//二叉树
type binaryTree struct {
	root   *BinTreeNode //根节点
	height int
	size   int
}

func NewBinaryTree(root *BinTreeNode) *binaryTree {
	return &binaryTree{root: root}
}


// 按层遍历
func (b *binaryTree) LevelOrder() *list.List{
	traversal := list.New()
	levelOrder(b.root, traversal)
	return traversal
}

func levelOrder(node *BinTreeNode, l *list.List){
	q := CreateQueue()
	q.Enqueue(node)

	for !q.isEmpty(){
		tmp := q.Dequeue().(*BinTreeNode)
		l.PushBack(tmp)
		if tmp.HasLChild(){
			q.Enqueue(tmp.lChild)
		}
		if tmp.HasRChild(){
			q.Enqueue(tmp.rChild)
		}
	}
}
