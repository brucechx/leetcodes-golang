package _63_二叉搜索树的第k结点

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
type binarySearchTree struct {
	root   *BinTreeNode //根节点
	height int
	size   int
}

func NewBinarySearchTree(root *BinTreeNode) *binarySearchTree {
	return &binarySearchTree{root: root}
}

func (b *binarySearchTree) kthNode(k int) *BinTreeNode{
	if b.root == nil{
		return nil
	}
	return kthNodeCore(b.root, k)
}

func kthNodeCore(root *BinTreeNode, k int) *BinTreeNode {
	//中序遍历
	s:=NewStack()
	GetBTVal(root ,s )
	for index,v:=range s.Val{
		if index+1==k{
			return v
		}

	}
	return nil
}

type Stack struct{
	Val []*BinTreeNode
}

func NewStack()*Stack{
	return &Stack{Val:make([]*BinTreeNode,0)}
}

func (s *Stack)Push(v *BinTreeNode){
	s.Val=append(s.Val, v)

}

func (s *Stack)Pop()*BinTreeNode{
	v:=s.Val[len(s.Val)-1]
	s.Val = s.Val[:len(s.Val) - 1]
	return v
}

func (s *Stack)IsEmpty()bool{
	return len(s.Val) == 0
}

func GetBTVal(root *BinTreeNode,s *Stack){
	if root==nil{
		return
	}
	GetBTVal(root.lChild,s)
	s.Val =append(s.Val,root)
	GetBTVal(root.rChild,s)
}