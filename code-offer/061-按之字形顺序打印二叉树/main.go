package _61_按之字形顺序打印二叉树

import "fmt"

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

func (b *binaryTree) PrintZ(){
	if b.root == nil{
		return
	}
	var current, reverse []*BinTreeNode
	flag := 0
	current = append(current, b.root)
	for len(current) > 0 {
		// 从最后一个开始取
		node := current[len(current) - 1]
		current = current[: len(current) - 1]
		fmt.Print(node.data)
		// 当前是从左往右打印的，那就按从左往右入栈
		if flag == 0{
			if node.lChild != nil{
				reverse = append(reverse, node.lChild)
			}
			if node.rChild != nil{
				reverse = append(reverse, node.rChild)
			}
		}else {
			// 当前是从右往左打印的，那就按从右往左入栈
			if node.rChild != nil{
				reverse = append(reverse, node.rChild)
			}
			if node.lChild != nil{
				reverse = append(reverse, node.lChild)
			}
		}
		if len(current) == 0{
			flag = 1 - flag
			current, reverse = reverse, current
			fmt.Println()
		}
	}
}
